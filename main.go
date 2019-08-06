package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brenoandrade/estrategia/api"
	"github.com/brenoandrade/estrategia/app"
)

func main() {
	app := app.New()

	app.StartServer()
	api.Init(app, app.Srv.Router)

	log.Println("[SERVER] running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
