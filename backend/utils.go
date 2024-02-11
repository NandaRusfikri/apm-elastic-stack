package backend

import (
	"context"
	"go.elastic.co/apm/v2"
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

func SendEmail(ctx context.Context, to string) error {
	span, ctx := apm.StartSpan(ctx, GetCurrentFunctionName(), "Utils")
	defer span.End()
	time.Sleep(time.Duration(RandInt(1000, 3000)) * time.Millisecond)
	return nil
}

func ExportPDF(ctx context.Context, data []string) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, "ExportPDF", GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(RandInt(1000, 3000)) * time.Millisecond)
	return data, nil
}
