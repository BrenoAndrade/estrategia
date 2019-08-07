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

// Handler estrutura para injetar as rotas do pacote mux
type Handler struct {
	router *mux.Router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// Server estrutura base do servidor
type Server struct {
	Router *mux.Router
	Server *http.Server
}

// App estrutura principal da aplicação
type App struct {
	Srv *Server
}

// StartServer inicia o servidor e os serviços
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
	services.InitRedis(config.URLRedis)

	log.Println("[SERVER] on:", app.Srv.Server.Addr)
	if err := app.Srv.Server.ListenAndServe(); err != nil {
		log.Println("[SERVER] off:", err)
	}
}

// New cria uma nova instancia do App
func New() *App {
	log.Println("[SERVER] initializing...")

	app := &App{
		Srv: &Server{
			Router: mux.NewRouter(),
		},
	}

	return app
}
