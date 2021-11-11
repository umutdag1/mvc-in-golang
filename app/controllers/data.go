package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/models"
	"github.com/rest-api/utils"
)

func GetAll(hp *utils.HttpPackage) {
	result, status, err := models.GetAllData()
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, err.Error()
	hp.SendResponse()
}

func Set(hp *utils.HttpPackage) {
	reqBody := models.Data{}
	if err := jsoner.DecodeJSON(hp.R.Body, &reqBody); err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, http.StatusInternalServerError, err.Error()
		hp.SendResponse()
		return
	}
	v := reflect.ValueOf(reqBody)
	structKeyValMap := make(map[string]interface{})
	for i := 0; i < v.Type().NumField(); i++ {
		structKeyValMap[v.Type().Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	if err := utils.StructHandler(structKeyValMap); err != nil {
		hp.Response.Result, hp.Response.Status, hp.Response.Error = nil, http.StatusBadRequest, err.Error()
		hp.SendResponse()
		return
	}
	result, status, err := models.AddData(&reqBody)
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, err.Error()
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
	hp.Response.Result, hp.Response.Status, hp.Response.Error = result, status, err.Error()
	hp.SendResponse()
}
