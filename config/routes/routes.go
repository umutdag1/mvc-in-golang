package routes

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/umutdag1/yemeksepeti-odev/app/controllers"
	"github.com/umutdag1/yemeksepeti-odev/app/libraries/logger"
	"github.com/umutdag1/yemeksepeti-odev/config"
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

func AuthRoute(handler, module interface{}) {
	if module == nil {
		logger.ErrorLogger.Println("Module Is Not Defined")
		os.Exit(3)
	}
	funcVal := reflect.ValueOf(handler)
	pointerOfFunc := funcVal.Pointer()
	runTimeOfFunc := runtime.FuncForPC(pointerOfFunc)
	curFilePath, _ := runTimeOfFunc.FileLine(pointerOfFunc)
	moduleType := reflect.TypeOf(module)
	fileNameFromModule := strings.ToLower(moduleType.Name()) + ".go"
	pkgPathPartials := strings.Split(moduleType.PkgPath(), "/")
	pkgPath := strings.Join(pkgPathPartials[len(pkgPathPartials)-2:], "/")
	expectedFilePath := ""
	if pkgPath == config.CONTROLLER_PATH {
		expectedFilePath = config.CUR_DIR + "/" + config.CONTROLLER_PATH + "/" + fileNameFromModule
	}
	if matched, _ := filepath.Match(expectedFilePath, curFilePath); !matched {
		logger.ErrorLogger.Printf("%q Module Or %q Function Is Not Defined", moduleType.Name(), runTimeOfFunc.Name())
		os.Exit(3)
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
