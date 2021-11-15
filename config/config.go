package config

import (
	"os"
	"strings"
	"time"
)

var (
	PROJECT_PATH    = "github.com/rest-api"
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
	DATA_JSON_FILE_NAME, DATA_JSON_FILE_EXT = FILE_TIME_STAMP_FUNC() + "-data", "json"
)
