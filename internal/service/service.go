package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService[T any] interface {
	getCollection() *mongo.Collection
	Create(data []byte) (T, error)
	Read(id primitive.ObjectID) (*T, error)
	Update(id primitive.ObjectID, newValues map[string]interface{}) (T, error)
	Delete(id primitive.ObjectID) error
}
