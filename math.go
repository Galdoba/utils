package utils

import (
	"fmt"
	"runtime"
)

//Min -
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Max -
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Assert -
func Assert(ok bool, s string) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("\033[31m\n%s:%d\n[error] %s\033[0m", file, line, s))
	}
}
