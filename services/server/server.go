package server

import (
	"log"
	"sja-boiler-plate/config"
	"sja-boiler-plate/routers"
)

func Init() {
	c := config.GetConfig()
	r := routers.NewRouter()
	err := r.Run(c.GetString("server.port"))
	if err != nil {
		log.Fatalln("server not started reason: ", err)
	}
}
