package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/Richtermnd/TestingSystem/application"
)

func main() {
	log := setupLogger()
	app := application.New(log)
	app.Init()
	go app.Run()

	// temporary solution
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info("Shutodwn")
}

func setupLogger() *slog.Logger {
	return slog.Default()
}
