package repository

import (
	"study/infra/database"
	"study/infra/repository/client"
	"study/infra/repository/login"
)

type Repositories struct {
	db     database.MongoDBInterFace
	Client client.RepositoryInterface // Interface, em vez de tipo concreto
	Login  login.RepositoryInterface  // Interface, em vez de tipo concreto
}

func NewRepositories(db database.MongoDBInterFace) *Repositories {
	clientRepository := client.NewRepositoryClient(db)
	loginRepository := login.NewRepositoryLogin(db)
	return &Repositories{
		db:     db,
		Client: clientRepository,
		Login:  loginRepository,
	}
}
