package xerr

type xErr struct {
	msg    string
	caller caller
	data   data
}

// New returns a xErr struct, that implements error interface.
func New(msg string) error {
	return xErr{
		msg:    msg,
		caller: call(),
		data:   make(data),
	}
}

func (e xErr) Error() string {
	return e.msg
}
