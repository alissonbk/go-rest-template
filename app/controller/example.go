package controller

import (
	"com.github.alissonbk/go-rest-template/app/exception"
	"com.github.alissonbk/go-rest-template/app/service"
	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	service *service.ExampleService
}

func NewExampleController(service *service.ExampleService) *ExampleController {
	return &ExampleController{service: service}
}

func (uc *ExampleController) GetAll(ctx *gin.Context) {
	defer exception.PanicHandler(ctx)
	ctx.JSON(200, uc.service.GetAll())
}
