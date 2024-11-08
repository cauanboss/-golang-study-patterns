package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDBAdapter struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoDBAdapter(uri string, dbName string, username string, password string) (*MongoDBAdapter, error) {
	client, err := createClient(uri, username, password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &MongoDBAdapter{
		client:   client,
		database: dbName,
	}, nil
}

func createClient(uri string, username string, password string) (*mongo.Client, error) {
	if username == "" || password == "" || uri == "" {
		panic("env not found")
	}

	clientOptions := options.Client().ApplyURI(uri)
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	clientOptions.SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Verifica a conexão
	if err := client.Ping(context.TODO(), nil); err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("Connected to MongoDB!")
	return client, nil
}

// GetCollection retorna uma coleção do banco de dados.
func (db *MongoDBAdapter) getCollection(collectionName string) *mongo.Collection {
	return db.client.Database(db.database).Collection(collectionName)
}

func (db *MongoDBAdapter) Disconnect() error {
	return db.client.Disconnect(context.TODO())
}

func (db *MongoDBAdapter) InsertOne(collectionName string, item interface{}) error {
	collection := db.getCollection(collectionName)
	fmt.Println("insert", collection, item)
	_, err := collection.InsertOne(context.TODO(), item)
	return err
}

func (db *MongoDBAdapter) FindOne(collectionName string, filter interface{}) (*mongo.SingleResult, error) {
	collection := db.getCollection(collectionName)
	return collection.FindOne(context.TODO(), filter), nil
}

func (db *MongoDBAdapter) DeleteOne(collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.getCollection(collectionName)
	return collection.DeleteOne(context.TODO(), filter)
}

func (db *MongoDBAdapter) FindAll(collectionName string) ([]interface{}, error) {
	collection := db.getCollection(collectionName)
	cursor, err := collection.Find(context.Background(), bson.D{}) // Usa um filtro vazio para retornar todos
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	var results []interface{}
	for cursor.Next(context.Background()) {
		var result map[string]interface{} // ou uma struct se conhecer o schema
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
