package app

import (
	"net/http"
	"strconv"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/services"
)

func getRecommendationTagsAndSave(repo *model.Repo, ch chan bool) {
	repo.Tags = services.GetWatsonKeywords(repo.URL)
	services.SetRepo(repo)
	ch <- true
}

// GetRepos pega todos os repos "starred" de um usuário
func (app *App) GetRepos(username string) ([]*model.Repo, *model.Error) {
	repos, err := services.GetRepos(username)
	if err != nil {
		return nil, err
	}

	size := len(repos)
	ch := make(chan bool, size)

	for _, repo := range repos {
		go getRecommendationTagsAndSave(repo, ch)
	}

	for i := 0; i < size; i++ {
		<-ch
	}

	return repos, nil
}

// AddTag adiciona uma tag a um repo
func (app *App) AddTag(repoID int, tag string) (*model.Repo, *model.Error) {
	repos, err := services.SearchRepos("*" + strconv.Itoa(repoID) + "*")
	if err != nil {
		return nil, err
	}

	var repo *model.Repo
	if len(repos) > 0 {
		repo = repos[0]
	}

	if repo == nil {
		return nil, model.NewError("app.repos.add_tag", "this repo not found", http.StatusNotFound)
	}

	key := repo.Key()
	if !repo.AddTag(tag) {
		return nil, model.NewError("app.repos.add_tag", "this tag has already been added", http.StatusConflict)
	}

	services.DelRepo(key)
	if err := services.SetRepo(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

// DelTag remove uma tag de um repo
func (app *App) DelTag(repoID int, tag string) (*model.Repo, *model.Error) {
	repos, err := services.SearchRepos("*" + strconv.Itoa(repoID) + "*")
	if err != nil {
		return nil, err
	}

	var repo *model.Repo
	if len(repos) > 0 {
		repo = repos[0]
	}

	if repo == nil {
		return nil, model.NewError("app.repos.add_tag", "this repo not found", http.StatusNotFound)
	}

	key := repo.Key()
	if !repo.DelTag(tag) {
		return nil, model.NewError("app.repos.del_tag", "tag not found", http.StatusNotFound)
	}

	services.DelRepo(key)
	if err := services.SetRepo(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

// Search procura entre as tags, repos que se encaixam na pesquisa
func (app *App) Search(pattern string) ([]*model.Repo, *model.Error) {
	return services.SearchRepos("*" + pattern + "*")
}
