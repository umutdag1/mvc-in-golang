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

func GetAll(w http.ResponseWriter, r *http.Request, resp *utils.ApiResponse) {
	result, status, err := models.GetAllData()
	resp.Result, resp.Status, resp.Error = result, status, err.Error()
	resp.SendResponse(w)
}

func Set(w http.ResponseWriter, r *http.Request, resp *utils.ApiResponse) {
	reqBody := models.Data{}
	if err := jsoner.DecodeJSON(r.Body, &reqBody); err != nil {
		resp.Result, resp.Status, resp.Error = nil, http.StatusInternalServerError, err.Error()
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
		err := fmt.Errorf("json: missing field %q", strings.Join(missings, ","))
		logger.ErrorLogger.Println(err.Error())
		resp.Result, resp.Status, resp.Error = nil, http.StatusBadRequest, err.Error()
		resp.SendResponse(w)
		return
	}
	result, status, err := models.AddData(&reqBody)
	resp.Result, resp.Status, resp.Error = result, status, err.Error()
	resp.SendResponse(w)
}

func Get(w http.ResponseWriter, r *http.Request, resp *utils.ApiResponse) {
	URIKey, err := utils.GetURIKeys(r, "key", 1)
	if err != nil {
		resp.Result, resp.Status, resp.Error = nil, http.StatusRequestedRangeNotSatisfiable, err.Error()
		resp.SendResponse(w)
		return
	}
	result, status, err := models.GetData(fmt.Sprintf("%v", URIKey.([]string)[0]))
	resp.Result, resp.Status, resp.Error = result, status, err.Error()
	resp.SendResponse(w)
}
