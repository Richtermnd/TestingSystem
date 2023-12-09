package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Richtermnd/TestingSystem/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

type Storage[T any] struct {
	log        *slog.Logger
	collection *mongo.Collection
}

func (s *Storage[T]) Create(newItem *T) (string, error) {
	s.log.Info("Creating item")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := s.collection.InsertOne(ctx, newItem)
	if err != nil {
		return "", storage.ErrBadInput
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	s.log.Info("Item created", "id", id.Hex())
	return id.Hex(), nil
}

func (s *Storage[T]) Read(pk string) (*T, error) {
	s.log.Info("Getting item", "pk", pk)
	if pk == "" {
		return nil, storage.ErrEmptyInput
	}

	s.log.Debug("Converting pk to ObjectId")
	id, err := primitive.ObjectIDFromHex(pk)
	if err != nil {
		return nil, storage.ErrBadInput
	}

	s.log.Debug("Getting items from database")
	var item T
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = s.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&item)
	if err != nil {
		return nil, storage.ErrNotFound
	}
	s.log.Info("Got item", "id", id.Hex())
	return &item, nil
}

func (s *Storage[T]) ReadAll() ([]*T, error) {
	s.log.Info("Getting all items")

	s.log.Debug("Getting cursor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, bson.D{})

	// idk what kind of error can be there, so not specific error for this
	// Later I will fix that
	if err != nil {
		return nil, err
	}

	s.log.Debug("Getting items from cursor")
	// Why `make`?
	// var items []*T serializing into null
	// items := make([]*T, 0) serializing into []
	items := make([]*T, 0)
	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}
	s.log.Info("Got all items")
	return items, nil
}

func (s *Storage[T]) Update(pk string, updatedItem *T) (string, error) {
	s.log.Info("Updating item", "pk", pk)
	if pk == "" {
		return "", storage.ErrEmptyInput
	}

	s.log.Debug("Converting pk to ObjectId")
	id, err := primitive.ObjectIDFromHex(pk)
	if err != nil {
		return "", storage.ErrBadInput
	}

	s.log.Debug("Updating")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := s.collection.ReplaceOne(ctx, bson.D{{"_id", id}}, updatedItem)
	if err != nil {
		return "", storage.ErrBadInput
	}

	// Update existing item or create new?
	if res.UpsertedCount == 0 {
		s.log.Info("Update existing item", "id", id)
		return id.Hex(), nil
	}
	upsertedId, _ := res.UpsertedID.(primitive.ObjectID)
	s.log.Info("Create new item", "id", upsertedId)
	return upsertedId.Hex(), nil
}

func (s *Storage[T]) Delete(pk string) (bool, error) {
	s.log.Info("Deleting item", "pk", pk)

	s.log.Debug("Converting pk to ObjectId")
	id, err := primitive.ObjectIDFromHex(pk)
	if err != nil {
		return false, err
	}

	s.log.Debug("Deleting")
	res, err := s.collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		return false, storage.ErrNotFound
	}
	s.log.Info("Item deleted")
	return res.DeletedCount == 1, nil
}

// NewStorage storage for one concrete model (T is model)
func NewStorage[T any](collectionName string, log *slog.Logger) *Storage[T] {
	collection := db.Collection(collectionName)
	return &Storage[T]{log: log, collection: collection}
}

type config struct {
	MongoURI string
	DBName   string
}

func Init(log *slog.Logger) {
	log.Info("Loading storage config.")
	cfg := loadConfig()
	log.Info("Connecting to mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// ? I really need it?
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		panic(err)
	}
	db = client.Database(cfg.DBName)
	log.Info("Succesfull connect")
}

func loadConfig() *config {
	mongoURI := mustGetEnv("MONGO_URI")
	dbName := mustGetEnv("MONGO_DB_NAME")
	return &config{MongoURI: mongoURI, DBName: dbName}
}

// mustGetEnv get variables from enviromnet, panic if not found.
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("empty %s enviroment variable", key))
	}
	return value
}
