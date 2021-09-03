package json

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Health struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	healthJson   = `{"status":"200","message":"Active!"}`
	healthStruct = &Health{
		Status:  "200",
		Message: "Active!",
	}
)

func TestUnmarshal(t *testing.T) {
	healthResult := &Health{}
	result := Unmarshal([]byte(healthJson), &healthResult)
	assert.Equal(t, healthStruct, healthResult)
	assert.Nil(t, result)
}

func TestMarshal(t *testing.T) {
	result, err := Marshal(healthStruct)

	assert.Equal(t, healthJson, string(result))
	assert.Nil(t, err)
}

func TestDecode(t *testing.T) {
	body := bytes.NewReader([]byte(`{"name":"test"}`))

	type DecodeStruct struct {
		Name string `json:"name"`
	}
	var bodyDecode DecodeStruct

	err := Decode(body, &bodyDecode)

	assert.Equal(t, "test", bodyDecode.Name)
	assert.Nil(t, err)

}
