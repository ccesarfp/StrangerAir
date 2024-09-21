package router

import "net/http"

type Route struct {
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func CreateRoutes() (routesGroups []map[string]map[string]Route) {
	routesGroups = []map[string]map[string]Route{}

	routesGroups = append(routesGroups, createAirQualityRoutes())

	return routesGroups
}
