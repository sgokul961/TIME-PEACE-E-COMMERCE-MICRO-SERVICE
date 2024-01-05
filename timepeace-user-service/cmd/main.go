package main

import (
	"log"

	"github.com/sgokul961/timepeace-user-service/pkg/config"
	"github.com/sgokul961/timepeace-user-service/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatalln("config err", configErr)
	}
	server, diErr := di.InitilizeApi(config)
	if diErr != nil {
		log.Fatal("cannot start server", diErr)
	} else {
		server.Start(config)
	}
}
