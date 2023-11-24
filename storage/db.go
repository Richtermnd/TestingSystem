package storage

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Config struct {
	MongoURI string
}

// Upload Config from enviroment
func loadConfig() (*Config, error) {
	// TODO: Load from env.
	return &Config{MongoURI: "mongodb://localhost:27017/"}, nil
}

// Init database module
func Init(log *slog.Logger) {
	// Load config
	log.Info("Load config")
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	// Connect to mongo
	log.Info("Connecting to mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client_, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		panic(err)
	}

	// Make client global for package
	log.Info("Succesfull connect.")
	client = client_
}

// Return database instance
func GetDB(dbName string) *mongo.Database {
	return client.Database(dbName)
}
