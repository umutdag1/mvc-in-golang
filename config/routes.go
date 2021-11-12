package config

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/rest-api/app/controllers"
	"github.com/rest-api/app/libraries/logger"
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

func AuthRoutes() {
	for _, route := range ROUTES {
		if route.Module == nil {
			logger.ErrorLogger.Println("Module Is Not Defined")
			os.Exit(3)
		}
		funcVal := reflect.ValueOf(route.Handler)
		pointerOfFunc := funcVal.Pointer()
		runTimeOfFunc := runtime.FuncForPC(pointerOfFunc)
		curFilePath, _ := runTimeOfFunc.FileLine(pointerOfFunc)
		moduleType := reflect.TypeOf(route.Module)
		fileNameFromModule := strings.ToLower(moduleType.Name()) + ".go"
		pkgPathPartials := strings.Split(moduleType.PkgPath(), "/")
		pkgPath := strings.Join(pkgPathPartials[len(pkgPathPartials)-2:], "/")
		expectedFilePath := ""
		if pkgPath == CONTROLLER_PATH {
			expectedFilePath = CUR_DIR + "/" + CONTROLLER_PATH + "/" + fileNameFromModule
		}
		if matched, _ := filepath.Match(expectedFilePath, curFilePath); !matched {
			logger.ErrorLogger.Printf("%q Module Or %q Function Is Not Defined", moduleType.Name(), runTimeOfFunc.Name())
			os.Exit(3)
		}
	}
}

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
