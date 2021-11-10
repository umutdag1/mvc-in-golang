package server

import (
	"net/http"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/config"
)

var (
	mux = http.NewServeMux()
)

func StartServer() {
	setHandlers()
	logger.InfoLogger.Println("Server is Started With Listening Port : " + config.API_PORT)
	logger.InfoLogger.Fatal(http.ListenAndServe(":"+config.API_PORT, mux))
}
