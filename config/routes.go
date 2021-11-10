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
	ROUTES = []Route{
		{
			Path:    "/api/getAll",
			Handler: controllers.GetAll,
			Method:  "GET",
		},
		{
			Path:    "/api/set",
			Handler: controllers.Set,
			Method:  "POST",
		},
		{
			Path:    "/api/get",
			Handler: controllers.Get,
			Method:  "GET",
		},
	}
)
