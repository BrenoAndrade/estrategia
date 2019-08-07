package app

import (
	"log"
	"net/http"
	"time"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/services"
	"github.com/gorilla/mux"
)

const (
	readTimeout  = 30
	writeTimeout = 30
)

// Handler for handler
type Handler struct {
	router *mux.Router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// Server struct http management
type Server struct {
	Router *mux.Router
	Server *http.Server
}

// App struct base application
type App struct {
	Srv *Server
}

// StartServer init listener
func (app *App) StartServer() {
	var handler http.Handler = &Handler{app.Srv.Router}

	config := model.GetConfig()

	app.Srv.Server = &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		Addr:         config.Port,
	}

	services.InitWatson(config.URLWatson, config.APIKey)

	log.Println("[SERVER] on:", app.Srv.Server.Addr)
	if err := app.Srv.Server.ListenAndServe(); err != nil {
		log.Println("[SERVER] off:", err)
	}
}

// New make instance App
func New() *App {
	log.Println("[SERVER] initializing...")

	app := &App{
		Srv: &Server{
			Router: mux.NewRouter(),
		},
	}

	return app
}
