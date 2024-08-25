package xerr

import (
	"encoding/json"
	"errors"
)

// JSON returns xErr struct as json-object with all fields, including data.
//
// If passed error is not an xErr instance, returns err.Error() result.
func JSON(err error) string {
	var e *xErr
	if !errors.As(err, &e) {
		return err.Error()
	}

	j, _ := json.Marshal(*e)

	return string(j)
}
