package app

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/brenoandrade/estrategia/model"
	"github.com/gorilla/mux"
	"github.com/tradersclub/TCIdeas/store"
)

const (
	readTimeout  = 10
	writeTimeout = 10
)

// Server struct http management
type Server struct {
	Store      store.Store
	Router     *mux.Router
	Server     *http.Server
	ListenAddr *net.TCPAddr
}

// StartServer init listener
func (app *App) StartServer() {
	app.Srv.Server = &http.Server{
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	addr := model.GetConfig().Port

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	app.Srv.ListenAddr = listener.Addr().(*net.TCPAddr)
	log.Println("[SERVER] on:", listener.Addr().String())

	go func() {
		var err error
		err = app.Srv.Server.Serve(listener)

		if err != nil && err != http.ErrServerClosed {
			log.Println("[SERVER] off:", err)
			time.Sleep(time.Second)
		}
	}()
}

// App struct base application
type App struct {
	Srv *Server
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
