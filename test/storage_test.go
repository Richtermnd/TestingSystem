package test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/Richtermnd/TestingSystem/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	testDB         = "testDB"
	testCollection = "testCollection"
)

type TestStruct struct {
	ValueString string
	ValueInt    int
	ValueSlice  []string
}

func TestInit(t *testing.T) {
	storage.Init(slog.Default(), testDB)
}

var collection *mongo.Collection

func TestGetDB(t *testing.T) {
	// Connect and drop collection.
	newCollection := storage.GetCollection(testCollection)
	err := newCollection.Drop(context.Background())
	if err != nil {
		t.Errorf("Error on getting database: %v", err)
		panic(err)
	}

	// Create empty collection.
	collection = newCollection
}

func TestGetEmptyCollection(t *testing.T) {
	// Getting cursor of all items.
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		t.Errorf("Error on getting cursor: %v", err)
	}

	var testStructs []TestStruct
	// Convert all items in structs.
	if err := cursor.All(context.Background(), &testStructs); err != nil {
		t.Errorf("Error on converting items in structs: %v", err)
	}

	if len(testStructs) != 0 {
		t.Errorf("Excepted 0 result, but got %d", len(testStructs))
	}
}
