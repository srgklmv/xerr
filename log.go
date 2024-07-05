package xerr

import (
	"encoding/json"
	"errors"
)

// Log returns xErr struct as json string with all fields including data.
func Log(err error) string {
	var e *xErr
	if !errors.As(err, &e) {
		return err.Error()
	}

	j, _ := json.Marshal(*e)

	return string(j)
}
