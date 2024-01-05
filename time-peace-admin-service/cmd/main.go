package main

import (
	"log"

	"github.com/sgokul961/time-peace-admin-service/pkg/config"
	"github.com/sgokul961/time-peace-admin-service/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatalln("failuire in loading config ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatalln("initilization error", diErr)
	} else {
		server.Start(config)

	}
}
