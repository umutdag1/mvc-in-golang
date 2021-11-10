package app

import (
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
)

func Init() {
	database.InitInMemDB()
	logger.InitLoggers()
}
