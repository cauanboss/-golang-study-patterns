package factory

import (
	"encoding/json"
	"fmt"
	"study/domain"
	"study/domain/dto"
	repository "study/infra/repository/client"
)

func (factory ClientFactory) InsertFactoryFromUseCase(dto dto.InsertDTO) repository.Client {
	return repository.Client{
		Name: dto.Name,
		Age:  dto.Age,
		Address: repository.Address{
			City:    dto.Address.City,
			Zipcode: dto.Address.Zipcode,
		},
	}
}

func (factory ClientFactory) InsertFactoryFromDB(data []interface{}) ([]domain.Client, error) {
	var clients []domain.Client

	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast item to map[string]interface{}")
		}

		jsonData, err := json.Marshal(itemMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item: %v", err)
		}

		var client domain.Client
		if err := json.Unmarshal(jsonData, &client); err != nil {
			return nil, fmt.Errorf("failed to unmarshal item: %v", err)
		}

		clients = append(clients, client)
	}

	return clients, nil
}
