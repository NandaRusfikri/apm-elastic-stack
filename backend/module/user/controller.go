package user

import (
	"backend"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Controller struct {
	gin     *gin.Engine
	usecase UsecaseInterface
}

func NewController(usecase UsecaseInterface, g *gin.Engine) {
	handler := &Controller{
		gin:     g,
		usecase: usecase,
	}
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
	g.POST("/forgot-password", handler.ForgotPassword)
	g.GET("/users", handler.GetList)
	g.PUT("/user", handler.Update)
	g.POST("/user/export-pdf", handler.ExportPDF)

}

func (ctrl *Controller) Login(ctx *gin.Context) {
	c := ctx.Request.Context()

	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()

	var request backend.RequestLogin

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.Login(c, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}

func (ctrl *Controller) Register(ctx *gin.Context) {
	c := ctx.Request.Context()
	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()
	var request backend.RegisterUser

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.Register(c, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}

func (ctrl *Controller) ForgotPassword(ctx *gin.Context) {
	c := ctx.Request.Context()
	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()
	var request backend.ForgotPassword

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.ForgotPassword(c, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}

func (ctrl *Controller) Update(ctx *gin.Context) {
	c := ctx.Request.Context()
	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()
	var request backend.User

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.usecase.Update(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}

func (ctrl *Controller) GetList(ctx *gin.Context) {
	c := ctx.Request.Context()
	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()

	data, err := ctrl.usecase.GetList(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}

func (ctrl *Controller) ExportPDF(ctx *gin.Context) {
	c := ctx.Request.Context()
	span, c := apm.StartSpan(c, backend.GetCurrentFunctionName(), "Controller")
	defer span.End()

	data, err := ctrl.usecase.ExportPDF(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success",
			"data": data})
		return
	}

}
