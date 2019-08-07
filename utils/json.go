package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ToJSON turn objects into json
func ToJSON(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		return make([]byte, 0)
	}
	return b
}

// ByteFromJSON make struct from json
func ByteFromJSON(b []byte, obj interface{}) {
	json.Unmarshal(b, &obj)
}

// ReaderFromJSON asd
func ReaderFromJSON(b io.Reader, obj interface{}) {
	bt, err := ioutil.ReadAll(b)
	if err == nil {
		ByteFromJSON(bt, &obj)
	}
}
