package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sja-boiler-plate/middlewares"
)

func ProtectedRoutes(r *gin.Engine) {
	r.GET("protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Protected Route")
	})
	r.Use(middlewares.AuthMiddleware())
}
