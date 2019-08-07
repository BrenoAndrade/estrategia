package api

import (
	"net/http"

	"github.com/brenoandrade/estrategia/app"
)

// Context struct para injetar o app e os parametros nos handlers
type Context struct {
	App    *app.App
	Params *Params
}

type handler struct {
	app        *app.App
	handleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handleFunc(&Context{
		App:    h.app,
		Params: apiParamsFromRequest(r),
	}, w, r)
}
