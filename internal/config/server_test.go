package config

import (
	"ccesarfp.com/StrangerAir/internal/enums/methods"
	"ccesarfp.com/StrangerAir/internal/routes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	t.Run("Shouldn't return nil when server is created using NewServer", func(t *testing.T) {
		server := NewServer()
		if server == nil {
			t.Error("Expected server to be non nil, but got nil")
		}
	})

	t.Run("Shouldn't return nil for mux when server is created using NewServer", func(t *testing.T) {
		server := NewServer()
		if server.mux == nil {
			t.Error("Expected mux to be non nil, but got nil")
		}
	})

	t.Run("Should have nil mux when server is created without using NewServer", func(t *testing.T) {
		server := Server{}
		if server.mux != nil {
			t.Error("Expected mux to be nil, but it is not")
		}
	})

	t.Run("Should have zero StartedAt when server is created without using NewServer", func(t *testing.T) {
		server := Server{}

		if !server.StartedAt.IsZero() {
			t.Error("Expected StartedAt to be zero, but it is not")
		}
	})
}

func TestServer_LoadEnv(t *testing.T) {
	server := NewServer()

	t.Run("Should load environment variables when APP_ENV_FILE is set to false", func(t *testing.T) {
		_ = os.Setenv("APP_ENV_FILE", "false")
		defer func() {
			_ = os.Unsetenv("APP_ENV_FILE")
		}()

		err := server.LoadEnv()
		if err != nil {
			t.Fatalf("Expected no error when loading env, but got: %s", err.Error())
		}
	})

	t.Run("Should set NAME to StrangerAir when APP_NAME is set to StrangerAir and APP_ENV_FILE is false", func(t *testing.T) {
		_ = os.Setenv("APP_ENV_FILE", "false")
		defer func() {
			_ = os.Unsetenv("APP_ENV_FILE")
		}()
		_ = os.Setenv("APP_NAME", "StrangerAir")
		defer func() {
			_ = os.Unsetenv("APP_NAME")
		}()

		_ = server.LoadEnv()

		if viper.GetString("NAME") != "StrangerAir" && viper.GetString("NAME") != "" {
			t.Fatalf("Expected environment variable NAME to be 'StrangerAir', but got: '%s'", viper.GetString("NAME"))
		}
	})

	t.Run("Shouldn't have NAME as empty when APP_NAME is set to StrangerAir", func(t *testing.T) {
		_ = os.Setenv("APP_ENV_FILE", "false")
		defer func() {
			_ = os.Unsetenv("APP_ENV_FILE")
		}()

		_ = os.Setenv("APP_NAME", "StrangerAir")
		defer func() {
			_ = os.Unsetenv("APP_NAME")
		}()

		_ = server.LoadEnv()

		if viper.GetString("NAME") == "" {
			t.Fatalf("Expected NAME to be non empty when APP_NAME is set, but got empty")
		}
	})

	t.Run("Shouldn't have NAME as empty when APP_NAME is set to StrangerAir", func(t *testing.T) {
		_ = os.Unsetenv("APP_NAME")
		if viper.GetString("NAME") != "" {
			t.Fatalf("Expected NAME to be empty when APP_NAME is not set, but got: '%s'", viper.GetString("NAME"))
		}
	})
}

func TestServer_RegisterRoutes(t *testing.T) {
	expected := "Testing"
	server := NewServer()

	t.Run("Should return correct status and body when a route is created and assigned", func(t *testing.T) {
		routeGroups := []map[string]map[string]router.Route{
			{
				"test": {
					"/get": router.Route{Method: methods.POST, Handler: func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusOK)
						_, _ = fmt.Fprintf(w, expected)
					}},
				},
			},
		}

		server.RegisterRoutes(routeGroups)

		req, err := http.NewRequest("POST", "/test/get", nil)
		if err != nil {
			t.Fatalf("Expected to be able to create the request, but got: %s", err.Error())
		}

		rr := httptest.NewRecorder()

		server.mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status code to be: %d, but got: %d", status, http.StatusOK)
		}

		if rr.Body.String() != expected {
			t.Errorf("Expected body to be: %s, but got: %s", rr.Body.String(), expected)
		}
	})

	t.Run("Should return error when route is duplicated", func(t *testing.T) {
		routeGroups := []map[string]map[string]router.Route{
			{
				"test": {
					"/get": router.Route{Method: methods.POST, Handler: func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusOK)
						_, _ = fmt.Fprintf(w, expected)
					}},
				},
			},
		}

		err := server.RegisterRoutes(routeGroups)
		err = server.RegisterRoutes(routeGroups)
		if err == nil {
			t.Errorf("Expected error when registering duplicated route, but got nothing")
		}
	})
}

func TestServer_Up(t *testing.T) {
	server := NewServer()

	t.Run("Should return an error when server port is not set and server cannot start", func(t *testing.T) {
		_ = server.LoadEnv()

		err := server.Up()
		if err == nil {
			t.Fatalf("Expected error to be non nil, but got nil")
		}
	})

	t.Run("Should return an error when server port is already in use", func(t *testing.T) {
		_ = os.Setenv("APP_ENV_FILE", "false")
		defer func() {
			_ = os.Unsetenv("APP_ENV_FILE")
		}()
		_ = os.Setenv("APP_SERVER_PORT", "8000")
		defer func() {
			_ = os.Unsetenv("APP_SERVER_PORT")
		}()
		_ = server.LoadEnv()

		go server.Up()

		time.Sleep(1 * time.Second)

		err := server.Up()
		if err == nil {
			t.Fatalf("Expected error to be non nil, but got nil")
		}
	})

}

func TestServer_GetLiteTime(t *testing.T) {
	server := NewServer()

	t.Run("Should return correct server lifetime when server is created", func(t *testing.T) {
		if server.GetLifeTime().Nanoseconds() == 0 {
			t.Fatalf("The lifetime was expected to be non zero, but got zero")
		}
	})

	t.Run("Should return zero lifetime when server is created without using NewServer", func(t *testing.T) {
		server = &Server{}

		if server.GetLifeTime() != 0 {
			t.Fatalf("The lifetime was expected to be zero, but got non zero")
		}
	})
}
