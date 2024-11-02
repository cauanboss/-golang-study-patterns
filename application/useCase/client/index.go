package client

import (
	"encoding/json"
	"study/infra/repository/client"
)

type UseCase struct {
	repository client.RepositoryInterface
}

func NewUseCaseClient(repository client.RepositoryInterface) *UseCase {
	return &UseCase{repository}
}

func (u *UseCase) Find(id string) ([]byte, error) {
	retFound, err := u.repository.Find("")
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(retFound)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (u *UseCase) InsertOne(dto client.InsertClient) (bool, error) {
	_, err := u.repository.InsertOne(dto)
	if err != nil {
		return false, err
	}
	return true, nil
}
