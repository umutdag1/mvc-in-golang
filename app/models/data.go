package models

import (
	"net/http"

	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/app/libraries/logger"
	"github.com/rest-api/database"
)

func GetAllData() (interface{}, int, error) {
	logger.InfoLogger.Println("Getting All Data")
	db := database.GetInMemDB()
	logger.InfoLogger.Println("Got All Data successfully")
	return db, http.StatusOK, nil
}

func GetData(key string) (interface{}, int, error) {
	logger.InfoLogger.Println("Getting Data")
	db := database.GetInMemDB()
	data, err := db.GetData(key)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return nil, http.StatusExpectationFailed, err
	}
	logger.InfoLogger.Println("Got Data Successfully")
	return data, http.StatusOK, nil
}

func AddData(reqBody *jsoner.Data) (interface{}, int, error) {
	logger.InfoLogger.Println("Adding Data")
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
	logger.InfoLogger.Println("Data Added Successfully")
	return data, http.StatusOK, nil
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
	return data, nil
}
