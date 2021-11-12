package app

import (
	"net/http"
	"reflect"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/config"
	"github.com/rest-api/utils"
)

func StartServer() {
	mux := http.NewServeMux()
	setHandlers(mux)
	logger.InfoLogger.Println("Server is Started With Listening Port : " + config.API_PORT)
	logger.InfoLogger.Fatal(http.ListenAndServe(":"+config.API_PORT, mux))
}

func setHandlers(mux *http.ServeMux) {
	for _, route := range config.ROUTES {
		expectedPath := config.PROJECT_PATH + "/" + config.CONTROLLER_PATH
		if packagePath := reflect.TypeOf(route.Module).PkgPath(); packagePath == expectedPath {
			handlerFunc := utils.CorsHandler(route.Handler, route.Method)
			mux.Handle(route.Path, handlerFunc)
		}
	}
}
