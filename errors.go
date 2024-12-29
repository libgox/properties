package properties

import "fmt"

// NoSuchPropertyError is returned when a property key value is missing.
type NoSuchPropertyError struct {
	Key string
}

func (e *NoSuchPropertyError) Error() string {
	return fmt.Sprintf("no such property '%s'", e.Key)
}

func NewNoSuchPropertyError(key string) *NoSuchPropertyError {
	return &NoSuchPropertyError{Key: key}
}
