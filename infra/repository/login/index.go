package login

import (
	"study/infra/database"
)

type IndexLogin struct {
	db             database.MongoDBInterFace
	collectionName string
}

func NewRepositoryLogin(db database.MongoDBInterFace) *IndexLogin {
	return &IndexLogin{db: db, collectionName: "login"}
}
