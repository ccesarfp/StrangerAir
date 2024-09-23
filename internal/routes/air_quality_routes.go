package router

import (
	"ccesarfp.com/StrangerAir/internal/enums/methods"
	"net/http"
)

// createAirQualityRoutes - create air quality routes and define handlers
func createAirQualityRoutes() (routeGroup map[string]map[string]Route) {
	const PREFIX string = "air-quality"
	routeGroup = map[string]map[string]Route{
		PREFIX: {
			"/get": Route{Method: methods.GET, Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}},
		},
	}

	return routeGroup
}
