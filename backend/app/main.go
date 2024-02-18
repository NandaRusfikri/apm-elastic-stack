package main

import (
	"backend"
	"backend/module/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/v2"
	"go.elastic.co/ecslogrus"
)

func init() {
	//log := log.New()
	log.SetFormatter(&ecslogrus.Formatter{})
	//log.SetOutput(os.Stdout)

	//file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	log.SetOutput(file)
	//} else {
	//	log.Info("Failed to log to file, using default stderr")
	//}
}

func main() {
	backend.TestConnectAPMServer()
	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}

	tracer, err := apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:        "Backend Nanda",
		ServiceVersion:     "V.1.2.3",
		ServiceEnvironment: "Misal DEV",
	})
	if err != nil {
		fmt.Println("err APM ", err)
	}

	opts := apmgin.WithTracer(tracer)
	apmgin.WithPanicPropagation()

	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine, opts))

	repo := user.NewRepository()
	service := user.NewUsecase(repo)

	user.NewController(service, engine)

	engine.Run(":9090")
}
