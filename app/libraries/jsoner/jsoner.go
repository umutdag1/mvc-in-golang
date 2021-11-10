package jsoner

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.ReadCloser, toAssign interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&toAssign)
	return err
}

func EncodeJSON(writer io.Writer, toSend interface{}) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(toSend)
	return err
}
