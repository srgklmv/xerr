package xerr

import (
	"runtime"
	"strings"
)

type caller string

func call() caller {
	_, f, l, _ := runtime.Caller(2)

	return caller(strings.Join([]string{f, string(l)}, ":"))
}
