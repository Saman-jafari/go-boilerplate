package server

import "sja-boiler-plate/config"

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.GetString("server.port"))
}
