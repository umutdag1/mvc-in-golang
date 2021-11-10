package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
)

type ApiResponse struct {
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
	Status int         `json:"status"`
}

func (response *ApiResponse) SendResponse(w http.ResponseWriter) {
	logger.InfoLogger.Println("sending response to client")
	w.WriteHeader(response.Status)
	if err := jsoner.EncodeJSON(w, response); err != nil {
		logger.ErrorLogger.Println(err.Error())
		w.Write([]byte(http.StatusText(response.Status)))
		w.WriteHeader(response.Status)
		return
	}
	logger.InfoLogger.Println("sent json to client successfully")
}

func GetURIKeys(r *http.Request, paramKey string, expectLen int) (interface{}, error) {
	URIKeys, OK := r.URL.Query()[paramKey]
	if !OK || len(URIKeys) != expectLen {
		errStr := "url param \"key\" is not existed or range is not satisfied"
		logger.ErrorLogger.Println(errStr)
		return nil, errors.New(errStr)
	}
	return URIKeys, nil
}

func CorsHandler(f http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			errStr := fmt.Sprintf("%v - %v Not Implemented", r.URL.Path, r.Method)
			logger.ErrorLogger.Printf(errStr)
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(errStr))
			return
		}
		logger.InfoLogger.Printf("%v - %v Received Request", r.URL.Path, r.Method)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		f(w, r)
	}
}
