package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// DI
	injection := NewInjection()
	exampleController := injection.NewExampleController()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "hello"})
	})
	api := router.Group("/api/v1")
	{
		example := api.Group("/example")
		example.GET("", exampleController.GetAll)
	}

	return router
}

// func Init() *echo.Echo {
// 	router := echo.New()
// 	router.Use(middleware.Logger())
// 	router.Use(middleware.Recover())

// 	// DI
// 	injection := NewInjection()
// 	exampleController := injection.NewExampleController()

// 	router.GET("", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
// 	})
// 	api := router.Group("/api/v1")

// 	api.GET("/ping", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "pong")
// 	})

// 	example := api.Group("/example")

// 	example.GET("", exampleController.GetAll)

// 	return router
// }

// func Init() *fiber.App {
// 	router := fiber.New()
// 	router.Use(logger.New())
// 	router.Use(recover.New())

// 	// DI
// 	injection := NewInjection()
// 	exampleController := injection.NewExampleController()

// 	router.Get("", func(c fiber.Ctx) error {
// 		c.Status(http.StatusOK)
// 		return c.JSON(fiber.Map{"message": "hello"})
// 	})
// 	api := router.Group("/api/v1")

// 	api.Get("/ping", func(c fiber.Ctx) error {
// 		c.Status(http.StatusOK)
// 		return c.SendString("pong")
// 	})

// 	example := api.Group("/example")

// 	example.Get("", exampleController.GetAll)

// 	return router
// }
