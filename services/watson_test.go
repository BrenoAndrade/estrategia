package services

import (
	"testing"

	"github.com/brenoandrade/estrategia/model"
)

var watsonError = "[ERROR] aguardando: %v, recebido: %v"

func TestInitWatson(t *testing.T) {
	config := model.GetConfig()
	InitWatson(config.URLWatson, config.APIKey)

	if watson == nil {
		t.Errorf(watsonError, "*NaturalLanguageUnderstandingV1", nil)
	}
}
