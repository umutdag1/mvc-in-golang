package app

import (
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
)

func Init() {
	logger.InitLoggers()
	database.InitInMemDB()
	//filer.CreateFile("a", "txt")
}
