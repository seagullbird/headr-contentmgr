package service

import (
	"fmt"
)

// ErrPathNotExist indicates a PathNotExist error
type ErrPathNotExist error

// ErrUnexpected indicates an unexpected error
type ErrUnexpected error

// MakeErrPathNotExist returns a PathNotExist error with the given path.
func MakeErrPathNotExist(path string) ErrPathNotExist {
	return fmt.Errorf("Path not exist: %s", path)
}

// MakeErrUnexpected returns all unexpected error with error message.
func MakeErrUnexpected(err error) ErrUnexpected {
	return fmt.Errorf("Unexpected error: %v", err)
}
