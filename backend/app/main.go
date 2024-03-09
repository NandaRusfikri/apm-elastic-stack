package main

import (
	"backend/module/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"os"
)

func init() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	os.MkdirAll("logs", 0777)

	file, err := os.OpenFile("logs/backend.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		logrus.Errorln("Error loading .env file ", err)
	}
	logrus.Infoln("✅ Success Load  .env file")

	// tracer, err := apm.NewTracerOptions(apm.TracerOptions{
	// 	ServiceName:        "Backend Nanda",
	// 	ServiceVersion:     "V.1.2.3",
	// 	ServiceEnvironment: "Misal DEV",
	// })
	// if err != nil {
	// 	fmt.Println("err APM ", err)
	// }

	//opts := apmgin.WithTracer(tracer)
	apmgin.WithPanicPropagation()

	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine))
	gin.SetMode(gin.ReleaseMode)
	repo := user.NewRepository()
	service := user.NewUsecase(repo)

	user.NewController(service, engine)

	logrus.Infoln("💜 Starting Backend Service :9090")

	engine.Run(":9090")
}
