package backend

import (
	"math/rand"
	"runtime"
	"time"
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

func SendEmail(to string) error {
	time.Sleep(time.Duration(RandInt(800, 1300)) * time.Millisecond)
	return nil
}
