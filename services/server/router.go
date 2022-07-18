package server

import (
	"github.com/gin-gonic/gin"
	"sja-boiler-plate/controllers/v1/health"
	"sja-boiler-plate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Use(middlewares.AuthMiddleware())

	//v1 := router.Group("v1")
	//{
	//	userGroup := v1.Group("user")
	//	{
	//		user := new(controllers.UserController)
	//		userGroup.GET("/:id", user.Retrieve)
	//	}
	//}
	return router

}
