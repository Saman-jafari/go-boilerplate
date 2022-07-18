package server

import (
	"log"
	"sja-boiler-plate/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	err := r.Run(c.GetString("server.port"))
	if err != nil {
		log.Fatalln("server not started reason: ", err)
	}
}
