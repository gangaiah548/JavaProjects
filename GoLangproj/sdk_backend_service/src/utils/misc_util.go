package utils

import (
	"runtime"
)

// getCurrentFuncName will return the current function's name.
// It can be used for a better log debug system.(I'm NOT sure.)
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
