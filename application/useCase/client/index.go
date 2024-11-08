package client

import (
	"encoding/json"
	"study/domain/dto"
	"study/domain/factory/client"
	"study/infra/repository/client"
)

type UseCase struct {
	repository    client.RepositoryInterface
	clientFactory factory.ClientFactory
}

func NewUseCaseClient(repository client.RepositoryInterface) *UseCase {
	return &UseCase{repository: repository, clientFactory: factory.ClientFactory{}}
}

func (u *UseCase) Find(id string) ([]byte, error) {
	retFound, err := u.repository.Find(id)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(retFound)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (u *UseCase) InsertOne(dto dto.InsertDTO) (bool, error) {
	insert := u.clientFactory.InsertFactoryFromUseCase(dto)
	_, err := u.repository.InsertOne(insert)
	if err != nil {
		return false, err
	}
	return true, nil
}
