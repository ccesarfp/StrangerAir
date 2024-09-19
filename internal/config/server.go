package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// Server - manages server start time and http configuration.
type Server struct {
	mux       *http.ServeMux
	StartedAt time.Time
}

// NewServer - creates a new instance of Server with the start time set to the current time.
func NewServer() *Server {
	return &Server{
		mux:       http.NewServeMux(),
		StartedAt: time.Now(),
	}
}

// LoadEnv - loads environment variables from the .env file and system environment variables.
func (s *Server) LoadEnv() error {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	if viper.GetString("ENV_FILE") != "false" {
		viper.SetConfigFile(".env")
		viper.AddConfigPath(".")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	return nil
}

// Up - start server and assign routes
func (s *Server) Up() error {
	err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")), s.mux)
	if err != nil {
		return err
	}
	return nil
}

// GetLifeTime - returns the duration of time elapsed since the server was started.
func (s *Server) GetLifeTime() time.Duration {
	return time.Since(s.StartedAt)
}
