package xerr

import "fmt"

type xErr struct {
	err    error
	Msg    string `json:"msg,omitempty"`
	Caller caller `json:"caller,omitempty"`
	Data   Data   `json:"data,omitempty"`
}

// New returns an xErr struct, that implements error interface.
func New(msg string) error {
	return &xErr{
		Msg:    msg,
		Caller: call(),
	}
}

// WrapError returns an xErr struct from err with new message.
func WrapError(err error, msg string) error {
	e := &xErr{
		err:    err,
		Caller: call(),
	}

	e.Msg = msg

	return e
}

func (e *xErr) Error() string {
	if e.err == nil {
		return e.Msg
	}

	return fmt.Sprintf("%s: %s", e.Msg, e.err)
}

func (e *xErr) Unwrap() error {
	return e.err
}
