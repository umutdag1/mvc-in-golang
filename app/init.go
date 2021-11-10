package app

import (
	"github.com/rest-api/app/libraries/databaser"
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/app/libraries/server"
)

func Init() {
	databaser.InitInMemDB()
	logger.InitLoggers()
	server.StartServer()
}
