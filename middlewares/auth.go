package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"sja-boiler-plate/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatusJSON(500, map[string]string{"msg": "something went wrong"})
			return
		}
		if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatusJSON(401, map[string]string{"msg": "unauthorized"})
			return
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatusJSON(401, map[string]string{"msg": "unauthorized"})
			return
		}
		c.Next()
	}
}
