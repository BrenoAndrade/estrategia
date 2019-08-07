package services

import (
	"testing"
)

var githubError = "[ERROR] aguardando: %v, recebido: %v"

func TestGetRepos(t *testing.T) {
	t.Parallel()

	value1, err1 := GetRepos("brenoandrade")
	if err1 != nil {
		t.Errorf(githubError, nil, err1)
	}

	if len(value1) == 0 {
		t.Errorf(githubError, len(value1), 0)
	}
}
