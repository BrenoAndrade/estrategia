package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Params struct padrão que contém todos os parâmetros declarados
type Params struct {
	RepoID   int
	Username string
	Tag      string
}

func apiParamsFromRequest(r *http.Request) *Params {
	params := &Params{}
	props := mux.Vars(r)

	params.Tag = props["tag"]
	params.Username = props["username"]

	if val, err := strconv.Atoi(props["repo_id"]); err == nil {
		params.RepoID = val
	}

	return params
}
