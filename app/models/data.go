package models

import (
	"net/http"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
	"github.com/rest-api/utils"
)

type Data struct {
	Key string      `json:"key"`
	Val interface{} `json:"value"`
}

func GetAllData() *utils.ApiResponse {
	logger.InfoLogger.Println("getting all data")
	resp := &utils.ApiResponse{}
	data, err := database.GetAllData()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		return resp
	}
	logger.InfoLogger.Println("got all data successfully")
	resp.Result = data
	resp.Status = http.StatusOK
	return resp
}

func GetData(key string) *utils.ApiResponse {
	logger.InfoLogger.Println("getting data")
	resp := &utils.ApiResponse{}
	data, err := database.GetData(key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		resp.Error = err.Error()
		resp.Status = http.StatusExpectationFailed
		return resp
	}
	logger.InfoLogger.Println("got data successfully")
	resp.Result = data
	resp.Status = http.StatusOK
	return resp
}

func AddData(reqBody *Data) *utils.ApiResponse {
	logger.InfoLogger.Println("adding data")
	err := database.AddData(reqBody.Key, reqBody.Val)
	resp := &utils.ApiResponse{}
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		resp.Error = err.Error()
		resp.Status = http.StatusInsufficientStorage
		return resp
	}
	data, err := database.GetData(reqBody.Key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		resp.Error = err.Error()
		resp.Status = http.StatusExpectationFailed
		return resp
	}
	logger.InfoLogger.Println("data added successfully")
	resp.Result = data
	resp.Status = http.StatusOK
	return resp
}
