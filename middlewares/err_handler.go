package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			if err != nil {
				log.Fatal(err)
			}
			// log, handle, etc.
		}
	}
}
