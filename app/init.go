package app

import (
	"time"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
	"github.com/rest-api/utils"
)

func Init() {
	logger.InitLoggers()
	database.InitInMemDB()
	go callSaveJSONDBFunc(10)
}

func callSaveJSONDBFunc(duration int64) {
	totalDuration := time.Second * time.Duration(duration)
	time.Sleep(totalDuration)
	logger.InfoLogger.Println("Automatic Calling Saving JSON DB File")
	utils.SaveJSONDBFile(database.GetInMemDB())
	logger.InfoLogger.Println("Automatic Calling Saving JSON DB File Successful")
	go callSaveJSONDBFunc(duration)
}
