package raftkv

import (
	"fmt"
)

type SnapshotStoreError struct {
	Path string
}

type RaftInstantiationError struct {
}

type InvalidLogEntryTypeError struct {
	LogEntryType interface{}
}

type InvalidKeyError struct {
	Key string
}

func (e *SnapshotStoreError) Error() string {
	return fmt.Sprintf("[RAFTKV]: Error creating snapshot store at '%v'.", e.Path)
}

func (e *RaftInstantiationError) Error() string {
	return fmt.Sprintf("[RAFTKV]: Could not instantiate a Raft object.")
}

func (e *InvalidLogEntryTypeError) Error() string {
	return fmt.Sprintf("[RAFTKV]: Invalid log entry type '%s'.", e.LogEntryType)
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("[RAFTKV]: Key '%s' is not in storage.", e.Key)
}
