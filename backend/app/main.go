package main

import (
	"backend"
	"backend/module/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/module/apmotel/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file ", err)
	}

	provider, err := apmotel.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	backend.Tracer = otel.Tracer("Service-nanda")

	r := gin.Default()
	r.Use(common())

	apmgin.WithPanicPropagation()
	r.Use(apmgin.Middleware(r))

	repo := auth.NewAuthRepository()
	service := auth.NewUsecase(repo)

	auth.NewController(service, r)

	r.Run(":9090")
}

func common() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("common %s %s\n", c.Request.Method, c.FullPath())
		ctx, span := backend.Tracer.Start(c.Request.Context(), fmt.Sprintf("%s %s", c.Request.Method, c.FullPath()))

		defer span.End()

		c.Set("OTEL", ctx)
		c.Next()

	}
}
