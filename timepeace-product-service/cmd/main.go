package main

import (
	"log"

	"github.com/sgokul961/timepeace-product-service/pkg/config"
	"github.com/sgokul961/timepeace-product-service/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatalln("failed to load config", configErr)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatalln("failed to initialize api", diErr)
	} else {
		server.Start(config)
	}

}
