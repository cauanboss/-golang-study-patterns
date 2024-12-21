package client

import (
	"study/domain"
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

func (u *UseCase) Find(id string) ([]domain.Client, error) {
	retFound, err := u.repository.Find(id)
	if err != nil {
		return nil, err
	}

	r, err := factory.ClientFactory{}.FindAllFactoryFromDB(retFound)

	if err != nil {
		return []domain.Client{}, err
	}
	return r, nil
}

func (u *UseCase) InsertOne(dto dto.InsertDTO) (bool, error) {
	insert := u.clientFactory.InsertFactoryFromUseCase(dto)
	_, err := u.repository.InsertOne(insert)
	if err != nil {
		return false, err
	}
	return true, nil
}
