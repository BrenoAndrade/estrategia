package services

import (
	"net/http"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/utils"
)

const base = "https://api.github.com"

// GetRepos aa
func GetRepos(username string) ([]*model.Repo, *model.Error) {
	url := utils.MakeURL(base, "users", username, "starred")

	res, err := http.Get(url)
	if err != nil {
		return nil, model.NewError("services.get_repos", err.Error(), res.StatusCode)
	}

	repos := make([]*model.Repo, 0)
	utils.ReaderFromJSON(res.Body, &repos)

	return repos, nil
}
