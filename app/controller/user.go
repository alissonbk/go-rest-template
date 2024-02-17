package controller

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/exception"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"com.github.alissonbk/go-rest-template/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	defer exception.PanicHandler(ctx)

	var userDTO dto.UserDTO
	err := ctx.BindJSON(&userDTO)
	if err != nil {
		log.Error(err)
		exception.PanicException(constant.ParsingFailed, "")
		return
	}

	save := uc.service.Save(userDTO.ToEntity())
	ctx.JSON(http.StatusOK, save)
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	return
}

func (uc *UserController) Update(ctx *gin.Context) {
	defer exception.PanicHandler(ctx)
	var userDTO dto.UserDTO
	err := ctx.BindJSON(&userDTO)
	if err != nil {
		log.Error(err)
		exception.PanicException(constant.ParsingFailed, "")
	}

	user := userDTO.ToEntity()
	user.Id, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Error(err)
		exception.PanicException(constant.InvalidRequest, "Invalid path Parameter")
	}
	uc.service.Update(user)
}

func (uc *UserController) Delete(ctx *gin.Context) {
	return
}
