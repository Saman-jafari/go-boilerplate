package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sja-boiler-plate/middlewares"
)

func ProtectedRoutes(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Protected Route")
	})
}
