package xerr

import (
	"fmt"
	"runtime"
)

type caller string

// call returns file and line, where error occurred.
func call() caller {
	_, f, l, _ := runtime.Caller(2)

	return caller(fmt.Sprintf("%s:%d", f, l))
}
