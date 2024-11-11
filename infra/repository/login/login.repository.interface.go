package login

type RepositoryInterface interface {
	Find(id string) (interface{}, error)
	InsertOne(dto Insert) (bool, error)
}
