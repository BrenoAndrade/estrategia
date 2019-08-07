package model

import "testing"

var errorError = "[ERROR] aguardando: %v, recebido: %v"

func TestNewError(t *testing.T) {
	t.Parallel()

	value := NewError("id", "message", 1)
	if value == nil {
		t.Errorf(errorError, "*model.Error", nil)
	}

	if value.ID != "id" {
		t.Errorf(errorError, "id", value.ID)
	}

	if value.Message != "message" {
		t.Errorf(errorError, "message", value.Message)
	}

	if value.Status != 1 {
		t.Errorf(errorError, 1, value.Status)
	}
}
