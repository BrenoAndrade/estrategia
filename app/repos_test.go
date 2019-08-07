package app

import (
	"testing"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/services"
)

var reposError = "[ERROR] aguardando: %v, recebido: %v"

func TestGetRepos(t *testing.T) {
	t.Parallel()

	config := model.GetConfig()
	services.InitRedis(config.URLRedis)
	services.InitWatson(config.URLWatson, config.APIKey)
	app := New()

	_, err := app.GetRepos("brenoandrade")
	if err != nil {
		t.Errorf(reposError, nil, "ERROR")
	}
}

func TestAddTag(t *testing.T) {
	t.Parallel()

	config := model.GetConfig()
	services.InitRedis(config.URLRedis)
	app := New()

	_, err := app.AddTag(123, "test1")
	if err == nil {
		t.Errorf(reposError, nil, "ERROR")
	}
}

func TestDelTag(t *testing.T) {
	t.Parallel()

	config := model.GetConfig()
	services.InitRedis(config.URLRedis)
	app := New()

	_, err := app.DelTag(123, "test1")
	if err == nil {
		t.Errorf(reposError, nil, "ERROR")
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	config := model.GetConfig()
	services.InitRedis(config.URLRedis)
	app := New()

	repos, _ := app.Search("*test1*")
	if len(repos) > 0 {
		t.Errorf(reposError, 0, len(repos))
	}
}
