package api

import (
	"net/http"

	"github.com/brenoandrade/estrategia/app"
	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/utils"
	"github.com/gorilla/mux"
)

// Routes struct das rotas default
type Routes struct {
	Root *mux.Router
}

// API struct principal do pacote
type API struct {
	App        *app.App
	BaseRoutes *Routes
}

// Public é um middleware para as rotas
func (api *API) Public(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:        api.App,
		handleFunc: h,
	}
}

// Init inicia o api (rotas)
func Init(a *app.App, root *mux.Router) *API {
	api := &API{
		App:        a,
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root

	api.initRepositories()

	root.HandleFunc("/status", returnStatusOK)
	root.HandleFunc("/{anything:.*}", Handle404)

	return api
}

// Handle404 rota padrão para notfound
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
