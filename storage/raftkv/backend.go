// Based on https://github.com/otoolep/hraftd/

package raftkv

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/raft"
	"github.com/mds796/CSGY9223-Final/storage/raftkv/raftkvpb"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type RaftKV struct {
	RaftID            string
	RaftDir           string
	RaftAddress       string
	SnapshotThreshold int
	Timeout           time.Duration
	KV                map[string][]byte // key-value storage
	Raft              *raft.Raft
	mutex             sync.Mutex
}

func CreateRaftKV(raftID string, address string) *RaftKV {
	raftDir := "./" + raftID
	os.MkdirAll(raftDir, 0700)

	return &RaftKV{
		RaftID:            raftID,
		RaftDir:           raftDir,
		RaftAddress:       address,
		SnapshotThreshold: 100,
		Timeout:           10 * time.Second,
		KV:                map[string][]byte{},
		mutex:             sync.Mutex{},
	}
}

func (r *RaftKV) Open(standaloneCluster bool) error {
	// Setup Raft
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(r.RaftID)

	// Setup Raft gRPC
	// TODO
	addr, err := net.ResolveTCPAddr("tcp", r.RaftAddress)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(r.RaftAddress, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Setup snapshot store
	snapshotStore, err := raft.NewFileSnapshotStore(r.RaftDir, r.SnapshotThreshold, os.Stderr)
	if err != nil {
		return &SnapshotStoreError{Path: r.RaftDir}
	}
	logStore := raft.NewInmemStore()
	stableStore := raft.NewInmemStore()

	raftInstance, err := raft.NewRaft(config, r, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return &RaftInstantiationError{}
	}
	r.Raft = raftInstance

	if standaloneCluster {
		clusterConfig := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		r.Raft.BootstrapCluster(clusterConfig)
	}

	return nil
}

// Configuration changes in the Raft cluster
func (r *RaftKV) Join(nodeID string, addr string) error {
	log.Printf("[RAFTKV] Received join request for remote node '%s' at '%s'", nodeID, addr)

	configFuture := r.Raft.GetConfiguration()
	err := configFuture.Error()
	if err != nil {
		log.Printf("[RAFTKV] Failed to get Raft configuration future: %v", err)
		return err
	}

	for _, server := range configFuture.Configuration().Servers {
		// Remove existing nodes from config
		if server.ID == raft.ServerID(nodeID) || server.Address == raft.ServerAddress(addr) {
			// Ignore if this config is already done
			if server.ID == raft.ServerID(nodeID) && server.Address == raft.ServerAddress(addr) {
				return nil
			}

			indexFuture := r.Raft.RemoveServer(server.ID, 0, 0)
			err := indexFuture.Error()
			if err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	// Add node to Raft config
	f := r.Raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	err = f.Error()
	if err != nil {
		return err
	}

	log.Printf("[RAFTKV] Node %s at %s joined Raft cluster", nodeID, addr)
	return nil
}

func (r *RaftKV) Get(key string) ([]byte, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if value, ok := r.KV[key]; ok {
		return value, nil
	}
	return []byte{}, &InvalidKeyError{Key: key}
}

func (r *RaftKV) Put(key string, value []byte) error {
	return r.ApplyLogEntry(raftkvpb.LogEntryType_PUT, key, value)
}

func (r *RaftKV) Delete(key string) error {
	return r.ApplyLogEntry(raftkvpb.LogEntryType_DEL, key, []byte{})
}

func (r *RaftKV) Iterate(namespace string) map[string][]byte {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	kv := map[string][]byte{}
	for k, v := range r.KV {
		if strings.HasPrefix(k, namespace) {
			kv[k] = v
		}
	}

	return kv
}

func (r *RaftKV) ApplyLogEntry(t raftkvpb.LogEntryType, key string, value []byte) error {
	logEntry := &raftkvpb.LogEntry{
		Type:  t,
		Key:   key,
		Value: value,
	}

	b, err := proto.Marshal(logEntry)
	if err != nil {
		return err
	}
	f := r.Raft.Apply(b, r.Timeout)
	return f.Error()
}

// Apply log entry to state-machine
func (r *RaftKV) Apply(l *raft.Log) interface{} {
	var entry raftkvpb.LogEntry

	if err := proto.Unmarshal(l.Data, &entry); err != nil {
		panic("failed to unmarshal log entry")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	switch entry.Type {
	case raftkvpb.LogEntryType_PUT:
		r.KV[entry.Key] = entry.Value
		return nil
	case raftkvpb.LogEntryType_DEL:
		_, ok := r.KV[entry.Key]
		if ok {
			delete(r.KV, entry.Key)
			return nil
		}
		return &InvalidKeyError{Key: entry.Key}
	default:
		return &InvalidLogEntryTypeError{LogEntryType: entry.Type}
	}
}

// Restore key-value store from a snapshot
func (r *RaftKV) Restore(rc io.ReadCloser) error {
	buffer := []byte{}
	_, err := io.ReadFull(rc, buffer)
	if err != nil {
		return err
	}

	snapshot := raftkvpb.Snapshot{}
	err = proto.Unmarshal(buffer, &snapshot)
	if err != nil {
		return err
	}

	// Set the state from the snapshot, no lock required according to
	// Hashicorp docs.
	r.KV = snapshot.Store.KV
	return nil
}

// Current state of the key-value store
func (r *RaftKV) Snapshot() (raft.FSMSnapshot, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return &Snapshot{Store: &raftkvpb.KeyValue{KV: r.Iterate("")}}, nil
}

type Snapshot raftkvpb.Snapshot

// Write snapshot to disk
func (s *Snapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := proto.Marshal(s.Store)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (s *Snapshot) Release() {}
