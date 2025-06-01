package main

import (
	"log"
	"newgo/config"
	"newgo/router"
)

func main() {
	config.InitConfig()
	r := router.SetRouter()
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
		return
	}
}
