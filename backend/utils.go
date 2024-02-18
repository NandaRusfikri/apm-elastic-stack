package backend

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
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

	log.Info("Email Send Success")
	return nil
}

func ExportPDF(ctx context.Context, data []string) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, "ExportPDF", GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(RandInt(1000, 3000)) * time.Millisecond)
	return data, nil
}

func TestConnectAPMServer() {
	resp, err := http.Get(os.Getenv("ELASTIC_APM_SERVER_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Baca response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error membaca response body:", err)
		return
	}

	// Cetak status code dan response body
	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Response body:", string(body))
}
