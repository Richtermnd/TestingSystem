package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService[T any] interface {
	getCollection() *mongo.Collection
	Create(T) (*primitive.ObjectID, error)
	ReadOne(*primitive.ObjectID) (*T, error)
	ReadAll() ([]*T, error)
	Update(*primitive.ObjectID, T) (*primitive.ObjectID, error)
	Delete(*primitive.ObjectID) (bool, error)
}
