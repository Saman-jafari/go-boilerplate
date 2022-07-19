package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sja-boiler-plate/controllers/v1/health"
)

func PublicRoutes(r *gin.Engine) {
	r.GET("public", func(c *gin.Context) {
		c.String(http.StatusOK, "Public Route")
	})
	r.GET("/health", new(controllers.HealthController).Status)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
