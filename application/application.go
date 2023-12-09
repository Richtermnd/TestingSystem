package application

import (
	"log/slog"

	"github.com/Richtermnd/TestingSystem/server"
	"github.com/Richtermnd/TestingSystem/storage/mongodb"
	"github.com/Richtermnd/TestingSystem/utils"
)

type Application struct {
	log    *slog.Logger
	server *server.Server
}

func New(log *slog.Logger) *Application {
	return &Application{log: log}
}

func (a *Application) Init() {
	utils.LoadEnviroment(a.log)
	// storage init
	mongodb.Init(a.log)
	// server init
	server := server.NewServer(a.log)
	server.Init()
	a.server = server
}

func (a *Application) Run() {
	// server start
	a.server.Run()
}
