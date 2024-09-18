package main

import (
	"ccesarfp.com/StrangerAir/internal/config"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	server := config.NewServer()

	err := server.LoadEnv()
	if err != nil {
		log.Fatalf("Error reading config file. Message Error: %s", err)
	}

	log.Printf("Application initialization took %s", server.GetLifeTime())
	log.Printf("Server started on port: %s", viper.GetString("SERVER_PORT"))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")), nil))
}
