package main

import (
	"ccesarfp.com/StrangerAir/internal/config"
	router "ccesarfp.com/StrangerAir/internal/routes"
	"github.com/spf13/viper"
	"log"
)

func main() {
	server := config.NewServer()

	err := server.RegisterRoutes(router.CreateRoutes())
	if err != nil {
		log.Panicf("Duplicate route. Message Error: %s\n", err)
	}

	err = server.LoadEnv()
	if err != nil {
		log.Panicf("Error reading config file. Message Error: %s\n", err)
	}

	log.Printf("Application initialization took %s", server.GetLifeTime())
	log.Printf("Server started on port: %s", viper.GetString("SERVER_PORT"))

	log.Panicln(server.Up())
}
