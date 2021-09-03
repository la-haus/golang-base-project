package json

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Decode(body io.Reader, input interface{}) error {
	return json.NewDecoder(body).Decode(input)
}
