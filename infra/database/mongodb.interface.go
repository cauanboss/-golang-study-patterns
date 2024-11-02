package database

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBInterFace interface {
	InsertOne(collectionName string, item interface{}) error
	Disconnect() error
	FindOne(collectionName string, filter interface{}) (*mongo.SingleResult, error)
	DeleteOne(collectionName string, filter interface{}) (*mongo.DeleteResult, error)
	FindAll(collectionName string) ([]interface{}, error)
}
