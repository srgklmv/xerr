package xerr

import "errors"

type Data map[string]interface{}

// AddData adds Data to xErr instance. Iterates over passed map and saves values
// by given keys.
//
// Affects only xErr types. Other error implementations keep unchanged.
func AddData(err error, d Data) (ok bool) {
	var e *xErr
	if !errors.As(err, &e) {
		return false
	}

	for k, v := range d {
		e.Data[k] = v
	}

	return true
}

// GetData returns Data from xErr type. Returns err.Error() result if error is
// not an xErr type.
func GetData(err error) Data {
	var e xErr
	if !errors.As(err, &e) {
		return nil
	}

	return e.Data
}
