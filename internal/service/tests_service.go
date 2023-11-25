package service

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/Richtermnd/TestingSystem/pkg/tests"
	"github.com/Richtermnd/TestingSystem/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestService struct {
	log *slog.Logger
}

func (t *TestService) getCollection() *mongo.Collection {
	return storage.GetCollection("tests")
}

func (t *TestService) Create(data []byte) (*primitive.ObjectID, error) {
	t.log.Info("Creating test")
	var newTest *tests.Test
	if err := json.Unmarshal(data, &newTest); err != nil {
		return nil, err
	}
	collection := t.getCollection()
	res, err := collection.InsertOne(context.TODO(), newTest)
	if err != nil {
		return nil, err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	// Not need it, but anyway.
	if !ok {
		return nil, errors.New("idk")
	}
	// s := id.String()
	t.log.Info("Test created.", "id", id.String())
	return &id, nil
}

func (t *TestService) ReadOne(id *primitive.ObjectID) (*tests.Test, error) {
	collection := t.getCollection()
	var test tests.Test
	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&test) // noqa
	if err != nil {
		return nil, err
	}
	return &test, nil
}

func (t *TestService) ReadAll(
	id primitive.ObjectID,
	offset, limit int,
	filters map[string]interface{},
) ([]*tests.Test, error) {
	collection := t.getCollection()
	var tests []*tests.Test
	var bsonFilters bson.D
	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	// Linter hates me if I use filters != nil...
	if filters == nil {
	} else {
		for k, v := range filters {
			bsonFilters = append(bsonFilters, bson.E{k, v})
		}
	}
	cursor, err := collection.Find(context.TODO(), bsonFilters, options)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), tests); err != nil {
		return nil, err
	}
	return tests, nil
}

func (t *TestService) Update(id *primitive.ObjectID, data []byte) (*primitive.ObjectID, error) {
	var newTest tests.Test
	err := json.Unmarshal(data, &newTest)
	if err != nil {
		return nil, err
	}
	collection := t.getCollection()
	res, err := collection.ReplaceOne(context.TODO(), bson.D{{"_id", id}}, newTest)
	// t.log.Info(slog.Attr("res", res.UpsertedID.String()))
	// log.Println()
	if err != nil {
		return nil, err
	}
	if res.UpsertedCount == 0 {
		return id, nil
	}
	upsertedId, ok := res.UpsertedID.(primitive.ObjectID)
	// Not need it, but anyway.
	if !ok {
		return nil, errors.New("idk")
	}
	return &upsertedId, nil
}

func (t *TestService) Delete(id *primitive.ObjectID) (bool, error) {
	collection := t.getCollection()
	res, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", id}})
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}

func New(log *slog.Logger) *TestService {
	return &TestService{log}
}
