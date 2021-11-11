package config

import (
	"github.com/rest-api/app/controllers"
	"github.com/rest-api/utils"
)

type Route struct {
	Path    string
	Handler func(*utils.HttpPackage)
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
