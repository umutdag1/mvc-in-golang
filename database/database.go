package database

import (
	"fmt"

	"github.com/rest-api/app/libraries/folderer"
	"github.com/rest-api/app/libraries/jsoner"
	"github.com/rest-api/config"
	"github.com/rest-api/utils"
)

type InMemDB map[string]interface{}

var (
	inMemDB InMemDB
)

func InitInMemDB() {
	inMemDB = InMemDB(make(map[string]interface{}))
	//utils.ReadJSONDBFile()
	files, err := folderer.OpenFolder(config.OUTPUT_PATH)
	if err != nil {
		return
	}
	filesNames := files.GetFilesPathsInFolder()
	lastSavedFileName := filesNames[len(filesNames)-1]
	data := []jsoner.Data{}
	err = utils.ReadJSONDBFile(lastSavedFileName, data)
	for _, dataStruct := range data {
		fmt.Println(dataStruct)
		/*if err = inMemDB.AddData(dataStruct.Key, dataStruct.Val); err != nil {
			logger.ErrorLogger.Println(err.Error())
		}*/
	}
	if err != nil {
		return
	}

}

func GetInMemDB() InMemDB {
	return inMemDB
}

func (inMemDB *InMemDB) GetData(key string) (InMemDB, error) {
	tempMemDB := InMemDB(make(map[string]interface{}))
	if !inMemDB.findKey(key) {
		return nil, fmt.Errorf("key %q is not existed", key)
	}
	tempMemDB[key] = (*inMemDB)[key]
	return tempMemDB, nil
}

func (inMemDB *InMemDB) AddData(key string, val interface{}) error {
	if inMemDB.findKey(key) {
		return fmt.Errorf("key %q is existed", key)
	}
	(*inMemDB)[key] = val
	return nil
}

func (inMemDB *InMemDB) DeleteData(key string) error {
	if !inMemDB.findKey(key) {
		return fmt.Errorf("key %q is not existed", key)
	}
	delete(*inMemDB, key)
	return nil
}

func (inMemDB *InMemDB) findKey(key string) bool {
	_, isExist := (*inMemDB)[key]
	return isExist
}
