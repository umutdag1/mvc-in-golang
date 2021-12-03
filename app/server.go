package app

import (
	"net/http"

	"github.com/umutdag1/yemeksepeti-odev/app/libraries/logger"
	"github.com/umutdag1/yemeksepeti-odev/config"
	"github.com/umutdag1/yemeksepeti-odev/config/routes"
	"github.com/umutdag1/yemeksepeti-odev/utils"
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
	sm.mux.HandleFunc("/", utils.NotFoundHandler(routes.MatchRouteWithURL))
	for _, route := range routes.GetRoutes() {
		routes.AuthRoute(route.Handler, route.Module)
		handlerFunc := utils.CorsHandler(route.Handler, route.Method)
		sm.mux.Handle(route.Path, handlerFunc)
	}
}
