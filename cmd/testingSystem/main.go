package main

import (
	"log/slog"

	"github.com/Richtermnd/TestingSystem/storage"
)

func main() {
	log := setupLogger()
	storage.Init(log)
}

func setupLogger() *slog.Logger {
	return slog.Default()
}
