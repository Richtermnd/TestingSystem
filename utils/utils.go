package utils

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnviroment(log *slog.Logger) {
	// Getting type of enviroment
	envType := os.Getenv("TESTING_SYSTEM_ENV")
	if envType == "" {
		envType = "development"
	}
	if envType != "production" {
		log.Warn(fmt.Sprintf("Using %s enviroment, not for production.", envType))
	}
	// loading enviroment variables.
	envFile := fmt.Sprintf(".env.%s", envType)
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}
}
