package utils

import (
	"encoding/json"
	"log"
)

// Marshal func
func Marshal(data interface{}) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// Unmarshal func
func Unmarshal(data []byte, res interface{}) {
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal(err)
	}
}
