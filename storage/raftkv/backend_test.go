package raftkv

import (
	"testing"
	"time"
)

func createRaftKV() *RaftKV {
	r := CreateRaftKV("node0", "localhost:1234")
	r.Open(true)

	// Wait to rig an election
	time.Sleep(3 * time.Second)
	return r
}

func TestRaftKV_PutAndGet(t *testing.T) {
	r := createRaftKV()

	k := "key"
	v := []byte("value")
	err := r.Put(k, v)
	if err != nil {
		t.Fatal(err)
	}
	value, err := r.Get(k)
	if err != nil {
		t.Fatal(err)
	}
	if string(value) != string(v) {
		t.Fatalf("Expected '%v', received '%v'", string(v), string(value))
	}
}

func TestRaftKV_GetInvalidKey(t *testing.T) {
	r := createRaftKV()

	k := "key"
	_, err := r.Get(k)
	if err == nil {
		t.Fatal("Error in GET call with invalid key")
	}
}

func TestRaftKV_Delete(t *testing.T) {
	r := createRaftKV()

	k := "key"
	v := []byte("value")
	err := r.Put(k, v)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Delete(k)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Get(k)
	if err == nil {
		t.Fatal("Error in GET call with invalid key")
	}
}

// func TestRaftKV_DeleteInvalidKey(t *testing.T) {
// 	r := createRaftKV()

// 	k := "key"
// 	err := r.Delete(k)
// 	if err == nil {
// 		t.Fatalf("Error in DELETE call with invalid key")
// 	}
// }

func TestRaftKV_IterateWithNamespace(t *testing.T) {
	r := createRaftKV()

	kvs := map[string][]byte{
		"ns1/key1": []byte("value1"),
		"ns1/key2": []byte("value2"),
		"ns1/key3": []byte("value3"),
	}

	for k, v := range kvs {
		r.Put(k, v)
	}
	r.Put("ns2/key1", []byte("not-value"))

	for k, v := range r.Iterate("ns1/") {
		if string(kvs[k]) != string(v) {
			t.Fatalf("Expected '%v', received '%v'.", string(kvs[k]), string(v))
		}
	}
}

func TestRaftKV_IterateWithoutNamespace(t *testing.T) {
	r := createRaftKV()

	kvs := map[string][]byte{
		"ns1/key1": []byte("value1"),
		"ns1/key2": []byte("value2"),
		"ns1/key3": []byte("value3"),
		"ns2/key1": []byte("not-value"),
	}

	for k, v := range kvs {
		r.Put(k, v)
	}

	for k, v := range r.Iterate("") {
		if string(kvs[k]) != string(v) {
			t.Fatalf("Expected '%v', received '%v'.", string(kvs[k]), string(v))
		}
	}
}
