package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ToJSON transforma um objeto em []byte
func ToJSON(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		return make([]byte, 0)
	}
	return b
}

// ByteFromJSON instancia uma estrutura através de um []byte
func ByteFromJSON(b []byte, obj interface{}) {
	json.Unmarshal(b, &obj)
}

// ReaderFromJSON instancia uma estrutura através de um io.Reader
func ReaderFromJSON(b io.Reader, obj interface{}) {
	bt, err := ioutil.ReadAll(b)
	if err == nil {
		ByteFromJSON(bt, &obj)
	}
}
