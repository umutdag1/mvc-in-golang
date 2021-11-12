package config

import (
	"os"
	"strings"
)

var (
	PROJECT_PATH    = "github.com/rest-api"
	CONTROLLER_PATH = "app/controllers"
	API_PORT        = "8080"
	CUR_DIR         = func() string {
		dir, _ := os.Getwd()
		return strings.ReplaceAll(dir, "\\", "/")
	}()
)
