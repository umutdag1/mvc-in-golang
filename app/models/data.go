package models

import (
	"fmt"
	"net/http"

	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
)

type Data struct {
	Key string      `json:"key"`
	Val interface{} `json:"value"`
}

func GetAllData() (interface{}, int, error) {
	logger.InfoLogger.Println("getting all data")
	data, err := database.GetAllData()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}
	logger.InfoLogger.Println("got all data successfully")
	return data, http.StatusOK, fmt.Errorf("")
}

func GetData(key string) (interface{}, int, error) {
	logger.InfoLogger.Println("getting data")
	data, err := database.GetData(key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusExpectationFailed, err
	}
	logger.InfoLogger.Println("got data successfully")
	return data, http.StatusOK, fmt.Errorf("")
}

func AddData(reqBody *Data) (interface{}, int, error) {
	logger.InfoLogger.Println("adding data")
	err := database.AddData(reqBody.Key, reqBody.Val)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusInsufficientStorage, err
	}
	data, err := database.GetData(reqBody.Key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusExpectationFailed, err
	}
	logger.InfoLogger.Println("data added successfully")
	return data, http.StatusOK, fmt.Errorf("")
}
