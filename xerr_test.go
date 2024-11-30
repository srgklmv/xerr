package xerr

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// xerr.go section

func TestXErr(t *testing.T) {
	type i interface {
		Error() string
		Unwrap() error
		Is(error) bool

		New(string) error
		Wrap(error, string) error

		AddData(Data)
	}
}

func TestXErr_New(t *testing.T) {
	var e error

	e = New("some error message")

	if e == nil {
		t.Error("New() should return a non-nil result.")
	}

	var err *xErr
	if !errors.As(e, &err) {
		t.Error("New() should return an xErr type.")
	}
}

func TestXErr_Error(t *testing.T) {
	msg := "some error message 1 2 3"
	outerMsg := "amogus"

	err := New(msg)

	if err == nil {
		t.Error("Error is nil.")
	}

	if err.Error() != msg {
		t.Error("Error() should return passed message.")
	}

	err = WrapError(err, outerMsg)
	if err.Error() != fmt.Sprintf("%s: %s", outerMsg, msg) {
		t.Error("Error() should return passed message and inner error msg.")
	}
}

// caller.go section

func TestCaller_call(t *testing.T) {
	err := New("error message")

	_, f, l, _ := runtime.Caller(0)
	l = l - 2
	c := fmt.Sprintf("%s:%d", f, l)

	var e *xErr
	if !errors.As(err, &e) {
		t.Error("Wrong error type.")
	}

	if string(e.Caller) != c {
		t.Error("Wrong caller.")
	}
}

// data.go section

func TestData_AddData(t *testing.T) {
	var (
		d1 = Data{
			"field1": "value1",
			"field2": "value2",
		}
		d2 = Data{
			"field3": "value3",
		}
		d12 = Data{
			"field1": "value1",
			"field2": "value2",
			"field3": "value3",
		}
	)

	err := New("some message")

	ok := AddData(err, d1)
	if !ok {
		t.Error("AddData() failed.")
	}

	var e *xErr
	if !errors.As(err, &e) {
		t.Error("Wrong error type.")
	}

	equal := reflect.DeepEqual(e.Data, d1)
	if !equal {
		t.Error("Wrong data #1.")
	}

	ok = AddData(err, d2)
	if !ok {
		t.Error("AddData() - 2 failed.")
	}

	var e2 *xErr
	if !errors.As(err, &e2) {
		t.Error("Wrong error type.")
	}

	equal = reflect.DeepEqual(e.Data, d12)
	if !equal {
		t.Error("Wrong data #2.")
	}
}

func TestData_GetData(t *testing.T) {
	err := New("some message")

	d := Data{
		"field1": "value1",
		"field2": "value2",
	}

	AddData(err, d)

	ed := GetData(err)

	if !reflect.DeepEqual(ed, d) {
		t.Error("Wrong data.")
	}

	d["field3"] = "value3"

	AddData(err, Data{
		"field3": "value3",
	})

	ed = GetData(err)

	if !reflect.DeepEqual(ed, d) {
		t.Error("Wrong data - 2.")
	}
}
