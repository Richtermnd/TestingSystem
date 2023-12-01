package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Richtermnd/TestingSystem/pkg/tests"
	"github.com/Richtermnd/TestingSystem/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestService struct {
	log *slog.Logger
}

func (t *TestService) getCollection() *mongo.Collection {
	return storage.GetCollection("tests")
}

func (t *TestService) Create(newTest tests.Test) (*primitive.ObjectID, error) {
	t.log.Info("Creating test")
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

func (t *TestService) ReadAll() ([]*tests.Test, error) {
	collection := t.getCollection()
	var tests []*tests.Test
	t.log.Info("Getting all tests")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &tests); err != nil {
		return nil, err
	}
	return tests, nil
}

func (t *TestService) Update(id *primitive.ObjectID, newTest tests.Test) (*primitive.ObjectID, error) {
	collection := t.getCollection()
	res, err := collection.ReplaceOne(context.TODO(), bson.D{{"_id", id}}, newTest)
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

func NewTestService(log *slog.Logger) *TestService {
	return &TestService{log}
}
