package xerr

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

//xerr.go section

func TestXErr_New(t *testing.T) {
	var e error

	e = New("some error message")

	if e == nil {
		t.Error("New() should return a non-nil result")
	}

	var err xErr
	if !errors.As(e, &err) {
		t.Error("New() should return an xErr type")
	}
}

func TestXErr_Error(t *testing.T) {
	msg := "some error message 1 2 3"

	err := New(msg)

	if err.Error() != msg {
		t.Error("Error() should return passed message")
	}
}

// json.go section

func TestXErr_JSON(t *testing.T) {
	msg := "some error message"
	err := New(msg)

	d := Data{
		"key1": "value1",
		"key2": "value2",
	}

	ok := AddData(err, d)
	if !ok {
		t.Error("AddData() should return true when adding data")
	}

	_, f, _, _ := runtime.Caller(0)
	l := 42

	expected := xErr{
		Msg:    msg,
		Caller: caller(fmt.Sprintf("%s:%d", f, l)),
	}

	AddData(expected, d)

	var j1, j2 interface{}

	e := json.Unmarshal(JSON(err), &j1)
	if e != nil {
		t.Error("Unmarshall error")
	}

	e = json.Unmarshal(JSON(expected), &j2)
	if e != nil {
		t.Error("Unmarshall error")
	}

	if !reflect.DeepEqual(err, expected) {
		t.Error("JSONs not equal")
	}
}

// caller.go section

func TestCaller_call(t *testing.T) {
	err := New("error message")

	_, f, _, _ := runtime.Caller(0)
	l := 11
	c := fmt.Sprintf("%s:%d", f, l)

	var e xErr
	if !errors.As(err, &e) {
		t.Error("Wrong error type")
	}

	if string(e.Caller) != c {
		t.Error("Wrong caller")
	}
}

// data.go section

func TestData_AddData(t *testing.T) {
	err := New("some message")

	d := Data{
		"field1": "value1",
		"field2": "value2",
	}

	ok := AddData(err, d)
	if !ok {
		t.Error("AddData() failed")
	}

	var e xErr
	if !errors.As(err, &e) {
		t.Error("Wrong error type")
	}

	if len(e.Data) != 2 {
		t.Error("Wrong data length")
	}

	v1, ok1 := e.Data["field1"]
	v2, ok2 := e.Data["field2"]

	if !ok1 || !ok2 {
		t.Error("Data not added")
	}

	if v1 != d["field1"] || v2 != d["field2"] {
		t.Error("Wrong data")
	}

	d2 := Data{
		"field3": "value3",
	}

	var e2 xErr
	if !errors.As(err, &e2) {
		t.Error("Wrong error type")
	}

	if len(e2.Data) != 3 {
		t.Error("Wrong data length")
	}

	v3, ok3 := e.Data["field3"]

	if !ok3 {
		t.Error("Data not added")
	}

	if v3 != d2["field3"] {
		t.Error("Wrong data")
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
		t.Error("Wrong data")
	}

	d["field3"] = "value3"

	AddData(err, Data{
		"field3": "value3",
	})

	ed = GetData(err)

	if !reflect.DeepEqual(ed, d) {
		t.Error("Wrong data - 2")
	}
}
