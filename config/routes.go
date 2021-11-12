package config

import (
	"github.com/rest-api/app/controllers"
)

type Route struct {
	Path    string
	Handler interface{}
	Module  interface{}
	Method  string
}

var (
	ROUTES = []Route{
		{
			Path:    "/api/getAll",
			Handler: controllers.GetAll,
			Module:  controllers.Data{},
			Method:  "GET",
		},
		{
			Path:    "/api/set",
			Handler: controllers.Set,
			Module:  controllers.Data{},
			Method:  "POST",
		},
		{
			Path:    "/api/get",
			Handler: controllers.Get,
			Module:  controllers.Data{},
			Method:  "GET",
		},
		{
			Path:    "/api/flushAll",
			Handler: controllers.FlushAll,
			Module:  controllers.Data{},
			Method:  "DELETE",
		},
	}
)

func GetRoutes() []Route {
	return ROUTES
}

func GetRoutesPath() []string {
	paths := make([]string, 0)
	for _, route := range ROUTES {
		paths = append(paths, route.Path)
	}
	return paths
}

func MatchRouteWithURL(urlPath string) bool {
	for _, route := range ROUTES {
		if route.Path == urlPath {
			return true
		}
	}
	return false
}
