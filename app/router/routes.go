package router

import (
	"net/http"

	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// DI
	injection := NewInjection()
	userController := injection.NewUserController()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.BuildResponse[any](constant.Success, "Hello"))
	})
	api := router.Group("/api/v1")
	{
		// The User domain it's only for example purpose...
		user := api.Group("/user")
		user.GET("", userController.GetAll)
		user.POST("", userController.Save)
		user.GET("/:userID", userController.GetByID)
		user.PUT("/:userID", userController.Update)
	}

	return router
}
