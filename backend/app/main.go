package main

import (
	"backend/module/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}

	r := gin.Default()
	//r.Use(common())

	tracer, err := apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:        "Backend Nanda",
		ServiceVersion:     "V.1.2.3",
		ServiceEnvironment: "Misal DEV",
	})

	opts := apmgin.WithTracer(tracer)

	apmgin.WithPanicPropagation()
	r.Use(apmgin.Middleware(r, opts))

	repo := user.NewRepository()
	service := user.NewUsecase(repo)

	user.NewController(service, r)

	r.Run(":9090")
}

func common() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Printf("common %s %s\n", c.Request.Method, c.FullPath())
		span, ctx := apm.StartSpan(c.Request.Context(), fmt.Sprintf("%s %s", c.Request.Method, c.FullPath()), "Comoon")

		defer span.End()

		c.Set("OTEL", ctx)
		c.Next()

	}
}
