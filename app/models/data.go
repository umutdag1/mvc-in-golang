package models

import (
	"net/http"

	"github.com/rest-api/app/libraries/databaser"
	"github.com/rest-api/app/libraries/responser"
)

type StrKeyInterfVal struct {
	Key string      `json:"key"`
	Val interface{} `json:"value"`
}

func GetAllData() *responser.ApiResponse {
	resp := &responser.ApiResponse{}
	resp.Result = databaser.GetInMemDB()
	resp.Status = http.StatusOK
	return resp
}

func SetValToKey(reqBody *StrKeyInterfVal) *responser.ApiResponse {
	databaser.SetValToKey(reqBody.Key, reqBody.Val)
	resp := &responser.ApiResponse{}
	resp.Result = databaser.GetValWithKeyByKey(reqBody.Key)
	resp.Status = http.StatusOK
	return resp
}
