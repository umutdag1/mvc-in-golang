package database

import (
	"fmt"
)

type InMemDB map[string]interface{}

var (
	inMemDB InMemDB
)

func InitInMemDB() {
	inMemDB = InMemDB(make(map[string]interface{}))
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
