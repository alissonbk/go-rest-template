package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func Init() *gin.Engine {
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())

// 	// DI
// 	injection := NewInjection()
// 	userController := injection.NewUserController()

// 	router.GET("", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, map[string]string{"message": "hello"})
// 	})
// 	api := router.Group("/api/v1")
// 	{
// 		// The User domain it's only for example purpose...
// 		user := api.Group("/user")
// 		user.GET("", userController.GetAll)
// 		user.POST("", userController.Save)
// 		user.GET("/:userID", userController.GetByID)
// 		user.PUT("/:userID", userController.Update)
// 	}

// 	return router
// }

func Init() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// DI
	injection := NewInjection()
	exampleController := injection.NewUserController()

	router.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
	})
	api := router.Group("/api/v1")

	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	example := api.Group("/example")

	example.GET("", exampleController.GetAll)

	return router
}
