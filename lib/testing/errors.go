package testing

import (
	"errors"
	"testing"
)

var (
	// ErrMethodUnimplemented is a stopgap to ensure testing coverage
	ErrMethodUnimplemented = errors.New("method is unimplemented")
)

func AssertMethodImplemented(t *testing.T, err error) bool {
	if err == ErrMethodUnimplemented {
		t.Error(err)
		t.FailNow()
		return false
	}
	return true
}
