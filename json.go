package xerr

import (
	"encoding/json"
	"errors"
	"fmt"
)

// JSON returns xErr struct as json-object with all fields, including data.
//
// If passed error is not an xErr instance, returns err.Error() result.
func JSON(err error) string {
	fmt.Print("\n REMOVE ME! ", "err: ", err, "\n")
	var e xErr
	if !errors.As(err, &e) {
		return err.Error()
	}

	j, _ := json.Marshal(e)

	return string(j)
}
