package storage

import (
	"fmt"
)

type InvalidKeyError struct {
	Key string
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("[STORAGE]: Key %s is not in storage.", e.Key)
}
