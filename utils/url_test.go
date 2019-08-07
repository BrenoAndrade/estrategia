package utils

import "testing"

var urlError = "[ERROR] aguardando: %v, recebido: %v"

func TestMakeURL(t *testing.T) {
	t.Parallel()

	waiting := `http://localhost:3001/repos/brenoandrade`
	value := MakeURL("http://localhost:3001", "repos", "brenoandrade")

	if value != waiting {
		t.Errorf(urlError, waiting, value)
	}
}
