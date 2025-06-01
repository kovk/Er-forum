package main

import (
	"newgo/config"
	"newgo/router"
)

func main() {
	config.InitConfig()
	r := router.SetRouter()
	r.Run(config.AppConfig.Server.Port)
}
