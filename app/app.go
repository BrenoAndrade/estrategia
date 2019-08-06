package app

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/store"
	"github.com/gorilla/mux"
)

const (
	readTimeout  = 10
	writeTimeout = 10
)

// Wrapper for handler
type Wrapper struct {
	router *mux.Router
}

func (wrapper *Wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrapper.router.ServeHTTP(w, r)
}

// Server struct http management
type Server struct {
	Store      store.Store
	Router     *mux.Router
	Server     *http.Server
	ListenAddr *net.TCPAddr
}

// App struct base application
type App struct {
	Srv *Server
}

// StartServer init listener
func (app *App) StartServer() {
	app.Srv.Server = &http.Server{
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	config := model.GetConfig()

	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		return
	}

	app.Srv.ListenAddr = listener.Addr().(*net.TCPAddr)
	log.Println("[SERVER] on:", listener.Addr().String())

	app.Srv.Store = store.NewSQLSupplier(config.ConnectionString)

	go func() {
		var err error
		err = app.Srv.Server.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			log.Println("[SERVER] off:", err)
			time.Sleep(time.Second)
		}
	}()
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
