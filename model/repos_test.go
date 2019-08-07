package model

import "testing"

var reposError = "[ERROR] aguardando: %v, recebido: %v"
var objectRepo = &Repo{
	ID:          123,
	Description: "test",
	Language:    "golang",
	Name:        "test-unity",
	Tags:        []string{"test", "unity"},
	URL:         "",
}

func TestKey(t *testing.T) {
	t.Parallel()

	waiting := "123|test|unity"
	value := objectRepo.Key()
	if value != waiting {
		t.Errorf(errorError, waiting, value)
	}
}

func TestDuplicatedKey(t *testing.T) {
	t.Parallel()

	waiting := true
	value := objectRepo.duplicatedTag("test")
	if value != waiting {
		t.Errorf(errorError, waiting, value)
	}
}

func TestAddTag(t *testing.T) {
	repo := *objectRepo

	repo.AddTag("test2")
	waiting := 3
	value := len(repo.Tags)
	if value != waiting {
		t.Errorf(errorError, waiting, value)
	}
}

func TestDelTag(t *testing.T) {
	repo := *objectRepo

	repo.DelTag("test")
	waiting := 1
	value := len(repo.Tags)
	if value != waiting {
		t.Errorf(errorError, waiting, value)
	}
}
