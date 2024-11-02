package client

import "study/infra/database"

type IndexClient struct {
	db             database.MongoDBInterFace
	collectionName string
}

func NewRepositoryClient(db database.MongoDBInterFace) *IndexClient {
	return &IndexClient{db: db, collectionName: "client"}
}
