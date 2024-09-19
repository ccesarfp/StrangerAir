package config

import (
	"github.com/spf13/viper"
	"time"
)

// Server manages the server's start time and environment variables.
type Server struct {
	StartedAt time.Time
}

// NewServer creates a new instance of Server with the start time set to the current time.
func NewServer() *Server {
	return &Server{
		StartedAt: time.Now(),
	}
}

// LoadEnv loads environment variables from the .env file and system environment variables.
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

// GetLifeTime returns the duration of time elapsed since the server was started.
func (s *Server) GetLifeTime() time.Duration {
	return time.Since(s.StartedAt)
}
