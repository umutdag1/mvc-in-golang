package config

import (
	"net/http"

	"github.com/rest-api/app/controllers"
)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

var (
	ROUTES = []*Route{
		{
			Path:    "/api",
			Handler: controllers.Home,
			Method:  "GET",
		},
		{
			Path:    "/api/set",
			Handler: controllers.Set,
			Method:  "POST",
		},
	}
)
