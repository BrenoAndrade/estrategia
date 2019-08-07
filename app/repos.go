package app

import (
	"net/http"
	"strconv"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/services"
)

func getRecommendationTags(repo *model.Repo, ch chan bool) {
	repo.Tags = services.GetWatsonKeywords(repo.URL)
	ch <- true
}

// GetRepositories adsa
func (app *App) GetRepositories(username string) ([]*model.Repo, *model.Error) {
	repos, err := services.GetRepos(username)
	if err != nil {
		return nil, err
	}

	size := len(repos)
	ch := make(chan bool, size)

	for _, repo := range repos {
		go getRecommendationTags(repo, ch)
	}

	for i := 0; i < size; i++ {
		<-ch
	}

	return repos, nil
}

// AddTag adsad
func (app *App) AddTag(repoID int, tag string) (*model.Repo, *model.Error) {
	repos, err := services.SearchRepos(strconv.Itoa(repoID))
	if err != nil {
		return nil, err
	}

	var repo *model.Repo
	if len(repos) > 0 {
		repo = repos[0]
	}

	if !repo.AddTag(tag) {
		return nil, model.NewError("app.repos.add_tag", "this tag has already been added", http.StatusConflict)
	}

	if err := services.SetRepo(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

// DelTag adsad
func (app *App) DelTag(repoID int, tag string) (*model.Repo, *model.Error) {
	repos, err := services.SearchRepos(strconv.Itoa(repoID))
	if err != nil {
		return nil, err
	}

	var repo *model.Repo
	if len(repos) > 0 {
		repo = repos[0]
	}

	if !repo.DelTag(tag) {
		return nil, model.NewError("app.repos.del_tag", "tag not found", http.StatusNotFound)
	}

	if err := services.SetRepo(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

// Search asdasd
func (app *App) Search(pattern string) ([]*model.Repo, *model.Error) {
	return services.SearchRepos(pattern)
}
