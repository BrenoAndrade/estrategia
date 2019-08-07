package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	defaultPage  = 0
	defaultLimit = 15
)

// Params adasdsa
type Params struct {
	Page     int
	Limit    int
	RepoID   int
	Username string
	Tag      string
}

// APIParamsFromRequest a
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
