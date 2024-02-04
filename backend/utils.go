package backend

import (
	"math/rand"
	"runtime"
)

func GetCurrentFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := 0
	for i := len(funcName) - 1; i >= 0; i-- {
		if funcName[i] == '/' {
			lastSlash = i
			break
		}
	}
	return funcName[lastSlash+1:]
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}
