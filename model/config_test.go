package model

import "testing"

var configError = "[ERROR] aguardando: %v, recebido: %v"

func TestGetConfig(t *testing.T) {
	t.Parallel()

	config := GetConfig()
	if config.Port == "" {
		t.Errorf(configError, "ConfigFile", nil)
	}

	if config.APIKey == "" {
		t.Errorf(configError, "ConfigFile", nil)
	}

	if config.URLWatson == "" {
		t.Errorf(configError, "ConfigFile", nil)
	}

	if config.URLRedis == "" {
		t.Errorf(configError, "ConfigFile", nil)
	}
}
