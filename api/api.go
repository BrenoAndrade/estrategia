package api

import (
	"net/http"

	"github.com/brenoandrade/estrategia/app"
	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/utils"
	"github.com/gorilla/mux"
)

// Routes struct of router
type Routes struct {
	Root    *mux.Router
	APIRoot *mux.Router
}

// API struct main for api
type API struct {
	App        *app.App
	BaseRoutes *Routes
}

// Init start the module
func Init(a *app.App, root *mux.Router) *API {
	api := &API{
		App:        a,
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root

	root.Handle("/status", http.HandlerFunc(returnStatusOK))
	root.Handle("/{anything:.*}", http.HandlerFunc(Handle404))

	return api
}

// Handle404 route base for not found
func Handle404(w http.ResponseWriter, r *http.Request) {
	err := model.NewError("handle.404", "Not found.", http.StatusNotFound)

	w.WriteHeader(err.Status)
	w.Write(utils.ToJSON(err))
}

func returnStatusOK(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["status"] = "OK"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.ToJSON(resp)))
}
