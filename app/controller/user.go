package controller

import (
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"com.github.alissonbk/go-rest-template/app/service"
	"com.github.alissonbk/go-rest-template/app/utils/util_response"
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
	var userDTO dto.UserDTO
	err := ctx.BindJSON(&userDTO)
	if err != nil {
		log.Error("Failed to bind userDTO, Error: ", err)
		ctx.JSON(util_response.InvalidJson())
		return
	}

	save, err := uc.service.Save(userDTO.ParseUserDTOToEntity())
	if err != nil {
		log.Error("Error saving user, Error:", err)
		ctx.JSON(util_response.InternalError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, save)
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	return
}

func (uc *UserController) Update(ctx *gin.Context) {
	var userDTO dto.UserDTO
	err := ctx.BindJSON(&userDTO)
	if err != nil {
		log.Error("Failed to bind userDTO, Error: ", err)
		ctx.JSON(util_response.InvalidJson())
		return
	}

	user := userDTO.ParseUserDTOToEntity()
	var id int
	id, err = strconv.Atoi(ctx.Param("userID"))
	log.Infoln("ID: ", id)
	user.Id = id
	err = uc.service.Update(user)
	if err != nil {
		log.Error("Failed to UPDATE user, Error: ", err)
		ctx.JSON(util_response.InternalError(err.Error()))
		return
	}
	return
}

func (uc *UserController) Delete(ctx *gin.Context) {
	return
}
