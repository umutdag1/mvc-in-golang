package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/rest-api/app/libraries/filer"
	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/config"
)

type ApiResponse struct {
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
	Status int         `json:"status"`
}

type HttpPackage struct {
	W        http.ResponseWriter
	R        *http.Request
	Response *ApiResponse
}

func (hp *HttpPackage) SendResponse() {
	logger.InfoLogger.Println("Sending Response To Client")
	//SaveJSONDBFile(hp.Response.Result)
	hp.W.WriteHeader(hp.Response.Status)
	if err := jsoner.EncodeJSON(hp.W, hp.Response); err != nil {
		logger.ErrorLogger.Println(err.Error())
		hp.W.WriteHeader(http.StatusInternalServerError)
		hp.W.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	logger.InfoLogger.Println("Sent Response To Client Successfully")
}

func ReadJSONDBFile(fileName string, toAssign interface{}) error {
	logger.InfoLogger.Println("Reading JSON DB File")
	targetPath := config.OUTPUT_PATH + "/" + fileName
	f, err := filer.OpenFile(targetPath)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	data, err := f.ReadFile()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	err = f.CloseFile()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	err = jsoner.JSONStructParseFromByteData(data, toAssign)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println("Reading JSON DB File Successful")
	return nil
}

func SaveJSONDBFile(data interface{}) error {
	logger.InfoLogger.Println("Saving JSON DB File")
	if data == nil {
		return fmt.Errorf("data cannot be saved to file")
	}
	targetFullFileName := fmt.Sprintf("%v.%v", config.DATA_JSON_FILE_NAME, config.DATA_JSON_FILE_EXT)
	targetPath := fmt.Sprintf("%v/%v", config.OUTPUT_PATH, targetFullFileName)
	ReadJSONDBFile(targetFullFileName, data)
	dataByte, err := jsoner.JSONParseToByteData(data)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	f, err := filer.OpenFile(targetPath)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	err = f.WriteFile(dataByte)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	err = f.CloseFile()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println("Saving JSON DB File Successful")
	return nil
}

func GetURIKeys(r *http.Request, paramKey string, expectLen int) (interface{}, error) {
	URIKeys, OK := r.URL.Query()[paramKey]
	if !OK || len(URIKeys) != expectLen {
		err := fmt.Errorf("url param \"key\" is not existed or range is not satisfied")
		logger.ErrorLogger.Println(err.Error())
		return nil, err
	}
	return URIKeys, nil
}

func JSONStructHandler(structVal reflect.Value) error {
	missings := []string{}
	structType := structVal.Type()
	for i := 0; i < structType.NumField(); i++ {
		key := structType.Field(i).Tag.Get("json")
		val := structVal.Field(i).Interface()
		if val == "" || val == nil {
			missings = append(missings, key)
		}
	}
	if len(missings) > 0 {
		err := fmt.Errorf("json: missing field %q", strings.Join(missings, ","))
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	return nil
}

func CorsHandler(f interface{}, method string) http.HandlerFunc {
	ControllerFunc, isControllerFuncExist := f.(func(*HttpPackage))
	return func(w http.ResponseWriter, r *http.Request) {
		hp := &HttpPackage{
			W:        w,
			R:        r,
			Response: &ApiResponse{},
		}
		if r.Method != method || !isControllerFuncExist {
			err := fmt.Errorf("%v - %v not implemented", r.URL.Path, r.Method)
			logger.ErrorLogger.Printf(err.Error())
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(err.Error()))
			return
		}
		logger.InfoLogger.Printf("%v - %v Received Request", r.URL.Path, r.Method)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ControllerFunc(hp)
	}
}

func NotFoundHandler(MatchRouteWithURLFunc func(string) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isURLPathExist := false
		if MatchRouteWithURLFunc(r.URL.Path) {
			isURLPathExist = true
		}

		if !isURLPathExist {
			err := fmt.Errorf("%v - %v not found", r.URL.Path, r.Method)
			logger.ErrorLogger.Printf(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
