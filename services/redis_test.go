package services

import (
	"testing"

	"github.com/brenoandrade/estrategia/model"
)

var redisError = "[ERROR] aguardando: %v, recebido: %v"
var objectRedis = &model.Repo{
	ID:          123,
	Description: "test",
	Language:    "golang",
	Name:        "test-unity",
	Tags:        []string{"test", "unity"},
	URL:         "",
}

func TestInitRedis(t *testing.T) {
	config := model.GetConfig()
	InitRedis(config.URLRedis)
	if redis == nil {
		t.Errorf(redisError, "*redis.Client", nil)
	}
}

func TestSetRepo(t *testing.T) {
	config := model.GetConfig()
	InitRedis(config.URLRedis)

	err := SetRepo(objectRedis)
	if err != nil {
		t.Errorf(redisError, nil, "ERROR")
	}
}

func TestSearchRepos(t *testing.T) {
	config := model.GetConfig()
	InitRedis(config.URLRedis)

	repos, err := SearchRepos("*123*")
	if err != nil {
		t.Errorf(redisError, nil, "ERROR")
	}

	if len(repos) == 0 {
		t.Errorf(redisError, 1, 0)
	}
}

func TestGetRepo(t *testing.T) {
	config := model.GetConfig()
	InitRedis(config.URLRedis)

	repo, err := GetRepo(objectRedis.Key())
	if err != nil {
		t.Errorf(redisError, nil, "ERROR")
	}

	if repo == nil {
		t.Errorf(redisError, objectRedis, nil)
	}
}

func TestDelRepo(t *testing.T) {
	config := model.GetConfig()
	InitRedis(config.URLRedis)

	DelRepo(objectRedis.Key())
	repo, err := GetRepo(objectRedis.Key())
	if err == nil {
		t.Errorf(redisError, "ERROR", nil)
	}

	if repo != nil {
		t.Errorf(redisError, nil, repo)
	}
}
