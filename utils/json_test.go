package utils

import (
	"strings"
	"testing"
)

var jsonError = "[ERROR] aguardando: %v, recebido: %v"

func TestToJSON(t *testing.T) {
	t.Parallel()

	waiting := `{"test":"unity"}`
	value := string(ToJSON(map[string]string{"test": "unity"}))

	if value != waiting {
		t.Errorf(jsonError, waiting, value)
	}
}

func TestByteFromJSON(t *testing.T) {
	t.Parallel()

	waiting := map[string]string{"test": "unity"}
	value := make(map[string]string)
	ByteFromJSON([]byte(`{"test":"unity"}`), &value)

	if value["test"] != waiting["test"] {
		t.Errorf(jsonError, waiting, value)
	}
}

func TestReaderFromJSON(t *testing.T) {
	t.Parallel()

	waiting := map[string]string{"test": "unity"}
	value := make(map[string]string)
	ReaderFromJSON(strings.NewReader(`{"test":"unity"}`), &value)

	if value["test"] != waiting["test"] {
		t.Errorf(jsonError, waiting, value)
	}
}
