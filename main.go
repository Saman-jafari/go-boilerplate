package main

import (
	"flag"
	"os"
	"sja-boiler-plate/services/server"

	//"articles/core/models"
	//"articles/routers/v1"
	"fmt"
	"sja-boiler-plate/config"
	"sja-boiler-plate/services/db/mysql"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	mysql.Init()
	fmt.Println("Server Running on Port: ", config.GetConfig().GetString("server.port"))
	server.Init()
}
