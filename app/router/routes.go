package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// DI
	injection := NewInjection()
	userController := injection.NewUserController()

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		user.GET("", userController.GetAll)
		user.POST("", userController.Save)
		user.GET("/:userID", userController.GetByID)
		user.PUT("/:userID", userController.Update)
		user.DELETE("/:userID", userController.Delete)

	}

	return router
}
