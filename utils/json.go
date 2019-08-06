package utils

import "encoding/json"

// ToJSON turn objects into json
func ToJSON(o interface{}) []byte {
	b, err := json.Marshal(o)
	if err != nil {
		return make([]byte, 0)
	}
	return b
}

// FromJSON make struct from json
func FromJSON(b []byte, o interface{}) {
	json.Unmarshal(b, &o)
}
