package jsoner

import (
	"encoding/json"
	"io"

	"github.com/umutdag1/yemeksepeti-odev/app/libraries/logger"
)

type Data struct {
	Key string      `json:"key"`
	Val interface{} `json:"value"`
}

func DecodeJSON(body io.ReadCloser, toAssign interface{}) error {
	logger.InfoLogger.Println("Decoding Body To JSON")
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&toAssign)
	if err != nil {
		return err
	}
	logger.InfoLogger.Println("Decoding Successful")
	return nil
}

func EncodeJSON(writer io.Writer, toSend interface{}) error {
	logger.InfoLogger.Println("Encoding Body To JSON")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(toSend)
	if err != nil {
		return err
	}
	logger.InfoLogger.Println("Encoding Successful")
	return nil
}

func JSONParseToByteData(data interface{}) ([]byte, error) {
	logger.InfoLogger.Println("Parsing JSON To Byte Data")
	byteData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return nil, err
	}
	logger.InfoLogger.Println("Parsing JSON To Byte Data Successful")
	return byteData, nil
}

func JSONStructParseFromByteData(data []byte, toAssign interface{}) error {
	logger.InfoLogger.Println("Parsing Byte Data To JSON")
	err := json.Unmarshal(data, &toAssign)
	if err != nil {
		return err
	}
	logger.InfoLogger.Println("Parsing Byte Data To JSON Successful")
	return nil
}
