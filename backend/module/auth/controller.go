package auth

import (
	"backend"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	gin     *gin.Engine
	usecase AuthUsecaseInterface
}

func NewController(usecase AuthUsecaseInterface, g *gin.Engine) {
	handler := &Controller{
		gin:     g,
		usecase: usecase,
	}
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
	g.POST("/forgot-password", handler.ForgotPassword)

}

func (ctrl *Controller) Login(ctx *gin.Context) {
	c := ctx.MustGet("OTEL").(context.Context)

	c, span := backend.Tracer.Start(c, backend.GetCurrentFunctionName())
	defer span.End()

	var request backend.RequestLogin

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.Login(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
	}

}

func (ctrl *Controller) Register(ctx *gin.Context) {
	c := ctx.Request.Context()
	c, span := backend.Tracer.Start(c, backend.GetCurrentFunctionName())
	defer span.End()
	var request backend.RegisterUser

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.Register(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
	}
	return

}

func (ctrl *Controller) ForgotPassword(ctx *gin.Context) {
	c := ctx.Request.Context()
	c, span := backend.Tracer.Start(c, backend.GetCurrentFunctionName())
	defer span.End()
	var request backend.ForgotPassword

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.ForgotPassword(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
	}
	return

}
