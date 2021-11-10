package controllers

import (
	"net/http"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/app/libraries/responser"
	"github.com/rest-api/app/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Getting All Data")
	resp := models.GetAllData()
	logger.InfoLogger.Println("Sending Response To Client")
	resp.SendResponse(w)
}

func Set(w http.ResponseWriter, r *http.Request) {
	logger.InfoLogger.Println("Setting Val To Key")
	reqBody := models.StrKeyInterfVal{}
	logger.InfoLogger.Println("Decoding Request Body To JSON")
	err := jsoner.DecodeJSON(r.Body, &reqBody)
	if err != nil {
		resp := &responser.ApiResponse{}
		resp.Error = err.Error()
		resp.Status = http.StatusRequestedRangeNotSatisfiable
		logger.ErrorLogger.Println(err.Error())
		resp.SendResponse(w)
		return
	}
	resp := models.SetValToKey(&reqBody)
	logger.InfoLogger.Println("Sending Requested-Body Response To Client")
	resp.SendResponse(w)
}
