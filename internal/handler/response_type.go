package handler

import (
	"encoding/json"
	"log"
)

type responseErrorWrapper struct {
	ErrMsg string `json:"error,omitempty"`
}

func (respErr responseErrorWrapper) jsonBytes() []byte {
	result, err := json.Marshal(respErr)
	if err != nil {
		log.Println("responseErrorWrapper.Marshal error: ", err.Error())
		return []byte(`{ "error": "unknown error" }`)
	}
	return result
}
