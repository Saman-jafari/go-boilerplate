package routers

import (
	"github.com/gin-gonic/gin"
	"sja-boiler-plate/middlewares"
	protected "sja-boiler-plate/routers/protected"
	public "sja-boiler-plate/routers/public"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORS())
	public.PublicRoutes(router)
	protected.ProtectedRoutes(router)

	return router
}
