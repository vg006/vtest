package app

import (
	"github.com/vg006/vtest/internal/logger"
	"github.com/vg006/vtest/internal/server"
	shut "github.com/vg006/vtest/internal/shutdown"
)

type App struct {
	logr *logger.Logger
	srv  *server.Server
}

func New(port string) *App {
	return &App{
		logr: logger.New(10),
		srv:  server.New(port),
	}
}

func (app *App) Exec() {
	shut.ListenForOSSignals()

	app.logr.Exec()
	app.srv.Exec()

	<-shut.ShutdownChan
	app.Exit()
}

func (app *App) Exit() {
	app.logr.LogInfo("Shutting down app...")
	app.srv.Exit()
	app.logr.LogDone("Shutting down logger...")
	app.logr.Exit("Gracefully fucking down app...")
}
