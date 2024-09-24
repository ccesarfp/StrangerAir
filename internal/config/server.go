package config

import (
	router "ccesarfp.com/StrangerAir/internal/routes"
	"errors"
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

// RegisterRoutes - register routes in the multiplexer
func (s *Server) RegisterRoutes(routeGroups []map[string]map[string]router.Route) {
	for _, routeGroup := range routeGroups {
		for prefix, routes := range routeGroup {
			for path, route := range routes {
				s.mux.HandleFunc(fmt.Sprintf("%s /%s%s", route.Method, prefix, path), route.Handler)
			}
		}
	}
}

// Up - start server and assign routes
func (s *Server) Up() error {
	serverPort := viper.GetString("SERVER_PORT")
	if serverPort == "" {
		return errors.New("server port not set")
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", serverPort), s.mux)
}

// GetLifeTime - returns the duration of time elapsed since the server was started.
func (s *Server) GetLifeTime() time.Duration {
	if s.StartedAt.IsZero() {
		return 0
	}
	return time.Since(s.StartedAt)
}
