package xerr

import "errors"

type data map[string]interface{}

// AddData affects only xErr error. Other implementations keeps unchanged.
// xErr uses Golang map as storage, so it behaves like usual map. Keep it in
// mind when use in concurrency or pass single key few times.
//
// Returns true if adding is successful, false if error is not an xErr struct.
func AddData(err error, key string, value interface{}) bool {
	var e *xErr
	if !errors.As(err, &e) {
		return false
	}

	e.Data[key] = value

	return true
}

// GetData returns Data by key from xErr struct. Returns nil if error is not an
// xErr struct or if key is not used.
func GetData(err error, key string) interface{} {
	var e *xErr
	if !errors.As(err, &e) {
		return nil
	}

	return e.Data[key]
}
