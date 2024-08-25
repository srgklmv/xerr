package test

import (
	"errors"
	"testing"
	"xerr"
)

func TestXErr_New(t *testing.T) {
	var e error

	e = xerr.New("some error message")

	if e == nil {
		t.Error("New() should return a non-nil result")
	}

	var err xerr.xErr
	if !errors.As(e, &err) {
		t.Error("New() should return an xErr type")
	}
}

func TestXErr_Error(t *testing.T) {
	msg := "some error message 1 2 3"

	err := xerr.New(msg)

	if err.Error() != msg {
		t.Error("Error() should return passed message")
	}
}
