package config

import (
	"github.com/rest-api/app/controllers"
	"github.com/rest-api/utils"
)

type Route struct {
	Path    string
	Handler func(*utils.HttpPackage)
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
	}
)
