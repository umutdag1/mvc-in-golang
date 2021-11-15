package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/models"
	"github.com/rest-api/utils"
)

type Data struct{}

func GetAll(hp *utils.HttpPackage) {
	result, status, err := models.GetAllData()
	if err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, status, err.Error()
		hp.SendResponse()
		return
	}
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, ""
	hp.SendResponse()
}

func Set(hp *utils.HttpPackage) {
	reqBody := jsoner.Data{}
	if err := jsoner.DecodeJSON(hp.R.Body, &reqBody); err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, http.StatusInternalServerError, err.Error()
		hp.SendResponse()
		return
	}
	if err := utils.StructHandler(reflect.ValueOf(reqBody)); err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, http.StatusBadRequest, err.Error()
		hp.SendResponse()
		return
	}
	result, status, err := models.AddData(&reqBody)
	if err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, status, err.Error()
		hp.SendResponse()
		return
	}
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, ""
	hp.SendResponse()
}

func Get(hp *utils.HttpPackage) {
	URIKey, err := utils.GetURIKeys(hp.R, "key", 1)
	if err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, http.StatusRequestedRangeNotSatisfiable, err.Error()
		hp.SendResponse()
		return
	}
	result, status, err := models.GetData(fmt.Sprintf("%v", URIKey.([]string)[0]))
	if err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, status, err.Error()
		hp.SendResponse()
		return
	}
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, ""
	hp.SendResponse()
}

func FlushAll(hp *utils.HttpPackage) {
	db := models.DeleteAllData()
	hp.Response.Result, hp.Response.Status, hp.Response.Error = db, http.StatusOK, fmt.Errorf("").Error()
	hp.SendResponse()
}
