package responser

import (
	"net/http"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
)

type ApiResponse struct {
	Error  string                 `json:"error"`
	Result map[string]interface{} `json:"result"`
	Status int                    `json:"status"`
}

func (response *ApiResponse) SendResponse(w http.ResponseWriter) {
	w.WriteHeader(response.Status)
	if err := jsoner.EncodeJSON(w, response); err != nil {
		w.Write([]byte(http.StatusText(response.Status)))
		w.WriteHeader(response.Status)
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Sent JSON To Client Successfully")
}
