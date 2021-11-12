package config

import (
	"os"
	"reflect"

	"github.com/rest-api/app/controllers"
	"github.com/rest-api/app/libraries/logger"
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

func GetRoutes() []Route {
	for _, route := range ROUTES {
		if route.Module == nil {
			logger.ErrorLogger.Println("module is not defined")
			os.Exit(3)
		}
		expectedPath := PROJECT_PATH + "/" + CONTROLLER_PATH
		if packagePath := reflect.TypeOf(route.Module).PkgPath(); packagePath != expectedPath {
			logger.ErrorLogger.Printf("%q module is not defined", reflect.TypeOf(route.Module).Name())
			os.Exit(3)
		}
	}
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
