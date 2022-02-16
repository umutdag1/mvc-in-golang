package app

import (
	"time"

	"github.com/umutdag1/mvc-in-golang/app/libraries/logger"
	"github.com/umutdag1/mvc-in-golang/config"
	"github.com/umutdag1/mvc-in-golang/database"
	"github.com/umutdag1/mvc-in-golang/utils"
)

func Init() {
	logger.InitLoggers()
	database.InitInMemDB()
	go callSaveJSONDBFunc(int64(config.DURATION_TIME_IN_SECONDS))
}

func callSaveJSONDBFunc(duration int64) {
	totalDuration := time.Second * time.Duration(duration)
	time.Sleep(totalDuration)
	logger.InfoLogger.Println("Automatic Calling Saving JSON DB File")
	utils.SaveJSONDBFile(database.GetInMemDB())
	logger.InfoLogger.Println("Automatic Calling Saving JSON DB File Successful")
	go callSaveJSONDBFunc(duration)
}
