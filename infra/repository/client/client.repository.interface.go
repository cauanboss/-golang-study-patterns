package client

type RepositoryInterface interface {
	Find(id string) (interface{}, error)
	InsertOne(dto InsertClient) (bool, error)
}
