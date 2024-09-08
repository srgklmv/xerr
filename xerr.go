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

func FromError(err error, msg string) error {

	return &xErr{
		err:    err,
		Msg:    msg,
		Caller: call(),
	}
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
