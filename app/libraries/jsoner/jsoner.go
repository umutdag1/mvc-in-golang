package jsoner

import (
	"encoding/json"
	"io"

	"github.com/rest-api/app/libraries/logger"
)

func DecodeJSON(body io.ReadCloser, toAssign interface{}) error {
	logger.InfoLogger.Println("decoding body to json")
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&toAssign)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println("decoding successful")
	return nil
}

func EncodeJSON(writer io.Writer, toSend interface{}) error {
	logger.InfoLogger.Println("encoding body to json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(toSend)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return err
	}
	logger.InfoLogger.Println("encoding successful")
	return nil
}
