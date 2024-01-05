package main

import (
	"log"

	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
	"github.com/sgokul961/timepeace-api-gateway/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatalln("configuration error:", configErr)
	}
	sever, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server ", diErr)
	} else {
		sever.Start(config)
	}
}
