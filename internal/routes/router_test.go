package router

import (
	"reflect"
	"testing"
)

func TestCreateRoutes(t *testing.T) {
	t.Run("Should be of type []map[string]map[string]router.Route ", func(t *testing.T) {
		routeGroups := CreateRoutes()
		if reflect.TypeOf(routeGroups).String() != "[]map[string]map[string]router.Route" {
			t.Fatalf("Expected []map[string]map[string]router.Route, got %s", reflect.TypeOf(routeGroups).String())
		}
	})
}
