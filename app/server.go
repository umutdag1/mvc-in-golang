package app

import (
	"net/http"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/config"
	"github.com/rest-api/utils"
)

type ServerMux struct {
	mux *http.ServeMux
}

func StartServer() {
	mux := ServerMux{mux: http.NewServeMux()}
	mux.setHandlers()
	logger.InfoLogger.Println("Server is Started With Listening Port : " + config.API_PORT)
	logger.InfoLogger.Fatal(http.ListenAndServe(":"+config.API_PORT, mux.mux))
}

func (sm ServerMux) setHandlers() {
	sm.mux.HandleFunc("/", utils.NotFoundHandler(config.MatchRouteWithURL))
	for _, route := range config.GetRoutes() {
		utils.AuthRoute(route.Handler, route.Module, config.CUR_DIR, config.CONTROLLER_PATH)
		handlerFunc := utils.CorsHandler(route.Handler, route.Method)
		sm.mux.Handle(route.Path, handlerFunc)
	}
}
