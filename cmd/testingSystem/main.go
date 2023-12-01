package main

import (
	"log/slog"

	"github.com/Richtermnd/TestingSystem/internal/application"
)

func main() {
	log := setupLogger()
	app := application.New(log)
	app.Run()
}

func setupLogger() *slog.Logger {
	return slog.Default()
}
