package storage

import (
	"fmt"
)

type ConnectionError struct {
}

type InvalidKeyError struct {
	Key string
}

type GetError struct {
	Key string
}

type PutError struct {
	Key string
}

type DeleteError struct {
	Key string
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("[STORAGE]: Could not connect to storage.")
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("[STORAGE]: Key '%s' is not in storage.", e.Key)
}

func (e *GetError) Error() string {
	return fmt.Sprintf("[STORAGE]: Getting key '%s' raised an error.", e.Key)
}

func (e *PutError) Error() string {
	return fmt.Sprintf("[STORAGE]: Putting key '%s' raised an error.", e.Key)
}

func (e *DeleteError) Error() string {
	return fmt.Sprintf("[STORAGE]: Deleting key '%s' raised an error.", e.Key)
}
