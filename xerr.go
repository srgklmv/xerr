package xerr

type xErr struct {
	Msg    string `json:"msg,omitempty"`
	Caller caller `json:"caller,omitempty"`
	Data   data   `json:"data,omitempty"`
}

// New returns an xErr struct, that implements error interface.
func New(msg string) error {
	return xErr{
		Msg:    msg,
		Caller: call(),
		Data:   make(data),
	}
}

func (e xErr) Error() string {
	return e.Msg
}
