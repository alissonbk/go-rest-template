package controller

import (
	"com.github.alissonbk/go-rest-template/app/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	return
}

func (uc *UserController) Save(ctx *gin.Context) {
	return
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	return
}

func (uc *UserController) Update(ctx *gin.Context) {
	return
}

func (uc *UserController) Delete(ctx *gin.Context) {
	return
}
