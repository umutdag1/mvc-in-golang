package database

import (
	"errors"
	"fmt"
)

type InMemDB map[string]interface{}

var (
	inMemDB InMemDB
)

func InitInMemDB() {
	inMemDB = InMemDB(make(map[string]interface{}))
}

func GetAllData() (InMemDB, error) {
	if inMemDB == nil {
		return nil, errors.New("inmemdb is not initiliazed")
	}
	return inMemDB, nil
}

func GetData(key string) (InMemDB, error) {
	tempMemDB := InMemDB(make(map[string]interface{}))
	if !findKey(key) {
		return nil, fmt.Errorf("key %q is not existed", key)
	}
	tempMemDB[key] = inMemDB[key]
	return tempMemDB, nil
}

func AddData(key string, val interface{}) error {
	if findKey(key) {
		return fmt.Errorf("key %q is existed", key)
	}
	inMemDB[key] = val
	return nil
}

func findKey(key string) bool {
	_, isExist := inMemDB[key]
	return isExist
}
