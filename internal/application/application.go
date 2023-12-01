package application

import (
	"context"
	"log/slog"

	"github.com/Richtermnd/TestingSystem/server"
	"github.com/Richtermnd/TestingSystem/storage"
)

type Application struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Application {
	return &Application{log: log}
}

func (a *Application) Run() {
	storage.Init(a.log, "testingSystem")
	server.Run(context.TODO(), a.log, ":8080")
}
