package storage

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

type Config struct {
	MongoURI string
	DBName   string
}

// Upload Config from enviroment
func loadConfig(dbName string) (*Config, error) {
	// TODO: Load from env.
	return &Config{
		MongoURI: "mongodb://localhost:27017/",
		DBName:   dbName,
	}, nil
}

// Init database module
func Init(log *slog.Logger, dbName string) {
	// Load config
	log.Info("Load config")
	cfg, err := loadConfig(dbName)
	if err != nil {
		panic(err)
	}

	// Connect to mongo
	log.Info("Connecting to mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		panic(err)
	}

	// Make client global for package
	log.Info("Succesfull connect.")
	db = client.Database(cfg.DBName)
}

// Return database instance
func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
