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
	db := database.GetInMemDB()
	logger.InfoLogger.Println("got all data successfully")
	return db, http.StatusOK, fmt.Errorf("")
}

func GetData(key string) (interface{}, int, error) {
	logger.InfoLogger.Println("getting data")
	db := database.GetInMemDB()
	data, err := db.GetData(key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusExpectationFailed, err
	}
	logger.InfoLogger.Println("got data successfully")
	return data, http.StatusOK, fmt.Errorf("")
}

func AddData(reqBody *Data) (interface{}, int, error) {
	logger.InfoLogger.Println("adding data")
	db := database.GetInMemDB()
	err := db.AddData(reqBody.Key, reqBody.Val)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusInsufficientStorage, err
	}
	data, err := db.GetData(reqBody.Key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusExpectationFailed, err
	}
	logger.InfoLogger.Println("data added successfully")
	return data, http.StatusOK, fmt.Errorf("")
}

func DeleteAllData() interface{} {
	db := database.GetInMemDB()
	for key, _ := range db {
		DeleteData(key)
	}
	return db
}

func DeleteData(key string) (interface{}, error) {
	db := database.GetInMemDB()
	data, err := db.GetData(key)
	if err != nil {
		return nil, err
	}
	err = db.DeleteData(key)
	if err != nil {
		return nil, err
	}
	return data, fmt.Errorf("")
}
