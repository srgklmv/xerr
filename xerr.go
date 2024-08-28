package xerr

type xErr struct {
	Msg    string `json:"msg,omitempty"`
	Caller caller `json:"caller,omitempty"`
	Data   Data   `json:"data,omitempty"`
}

// New returns an xErr struct, that implements error interface.
func New(msg string) error {
	return xErr{
		Msg:    msg,
		Caller: call(),
		Data:   make(Data),
	}
}

func (e xErr) Error() string {
	return e.Msg
}
