package config

import (
	"os"
	"strings"
	"time"
)

var (
	PROJECT_PATH    = "github.com/umutdag1/mvc-in-golang"
	CONTROLLER_PATH = "app/controllers"
	API_PORT        = "8080"
	CUR_DIR         = func() string {
		dir, _ := os.Getwd()
		return strings.ReplaceAll(dir, "\\", "/")
	}()
	OUTPUT_PATH          = CUR_DIR + "/tmp"
	FILE_TIME_STAMP_FUNC = func() string {
		timeStamp := time.Now().Format("2006-01-02T15:04:05")
		return strings.ReplaceAll(timeStamp, ":", "-")
	}
	DATA_JSON_FILE_NAME, DATA_JSON_FILE_EXT, DURATION_TIME_IN_SECONDS = FILE_TIME_STAMP_FUNC() + "-data", "json", 10
)
