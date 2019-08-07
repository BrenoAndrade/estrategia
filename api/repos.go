package api

import (
	"net/http"

	"github.com/brenoandrade/estrategia/utils"
)

func (api *API) initRepositories() {
	route := api.BaseRoutes.Root.Handle
	public := api.Public

	route("/v1/repos/{username:.*}/github", public(getRepositoriesFromGithub)).Methods("GET")
	route("/v1/repos/{repo_id:[0-9]+}/tags/{tag:.*}", public(addTag)).Methods("POST")
	route("/v1/repos/{repo_id:[0-9]+}/tags/{tag:.*}", public(delTag)).Methods("DELETE")
	route("/v1/tags/{tag:.*}", public(search)).Methods("GET")
}

func getRepositoriesFromGithub(c *Context, w http.ResponseWriter, r *http.Request) {
	if repos, err := c.App.GetRepositories(c.Params.Username); err != nil {
		w.WriteHeader(err.Status)
		w.Write(utils.ToJSON(err))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ToJSON(repos))
	}
}

func addTag(c *Context, w http.ResponseWriter, r *http.Request) {
	if repos, err := c.App.AddTag(c.Params.RepoID, c.Params.Username); err != nil {
		w.WriteHeader(err.Status)
		w.Write(utils.ToJSON(err))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(utils.ToJSON(repos))
	}
}

func delTag(c *Context, w http.ResponseWriter, r *http.Request) {
	if repos, err := c.App.DelTag(c.Params.RepoID, c.Params.Username); err != nil {
		w.WriteHeader(err.Status)
		w.Write(utils.ToJSON(err))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ToJSON(repos))
	}
}

func search(c *Context, w http.ResponseWriter, r *http.Request) {
	if repos, err := c.App.Search(c.Params.Tag); err != nil {
		w.WriteHeader(err.Status)
		w.Write(utils.ToJSON(err))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ToJSON(repos))
	}
}
