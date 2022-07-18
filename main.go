package main

import (
	"flag"
	"os"
	"sja-boiler-plate/services/server"

	//"articles/core/models"
	//"articles/routers/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"sja-boiler-plate/config"
	"sja-boiler-plate/services/db/mysql"
)

var router *gin.Engine

func init() {
	//mysql.GetDB()
	//router = gin.New()
	//router.NoRoute(noRouteHandler())
	//version1 := router.Group("/v1")
	//v1.InitRoutes(version1)
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	mysql.Init()
	fmt.Println("Server Running on Port: ", 9090)
	c := config.GetConfig()
	fmt.Println(c.GetString("db.host"))
	server.Init()

}
