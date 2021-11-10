package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/app/models"
	"github.com/rest-api/utils"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	resp := models.GetAllData()
	resp.SendResponse(w)
}

func Set(w http.ResponseWriter, r *http.Request) {
	reqBody := models.Data{}
	resp := &utils.ApiResponse{}
	if err := jsoner.DecodeJSON(r.Body, &reqBody); err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		resp.SendResponse(w)
		return
	}
	missings := []string{}
	if reqBody.Key == "" {
		missings = append(missings, "key")
	}
	if reqBody.Val == nil {
		missings = append(missings, "value")
	}
	if len(missings) > 0 {
		logger.ErrorLogger.Printf("json: missing field %q", strings.Join(missings, ","))
		resp.Error = fmt.Sprintf("json: missing field %q", strings.Join(missings, ","))
		resp.Status = http.StatusBadRequest
		resp.SendResponse(w)
		return
	}
	resp = models.AddData(&reqBody)
	resp.SendResponse(w)
}

func Get(w http.ResponseWriter, r *http.Request) {
	URIKey, err := utils.GetURIKeys(r, "key", 1)
	resp := &utils.ApiResponse{}
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusRequestedRangeNotSatisfiable
		resp.SendResponse(w)
		return
	}
	resp = models.GetData(fmt.Sprintf("%v", URIKey.([]string)[0]))
	resp.SendResponse(w)
}
