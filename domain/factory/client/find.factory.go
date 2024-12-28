package factory

import (
	"encoding/json"
	"fmt"
	"study/domain"
)

func (factory ClientFactory) FindFactoryFromDB(data map[string]interface{}) (domain.Client, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return domain.Client{}, fmt.Errorf("failed to marshal item: %v", err)
	}
	var client domain.Client
	if err := json.Unmarshal(jsonData, &client); err != nil {
		return domain.Client{}, fmt.Errorf("failed to unmarshal item: %v", err)
	}
	return client, nil
}

func (factory ClientFactory) FindAllFactoryFromDB(data []interface{}) ([]domain.Client, error) {
	var clients []domain.Client
	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return []domain.Client{}, fmt.Errorf("failed to cast item to map[string]interface{}")
		}
		client, err := factory.FindFactoryFromDB(itemMap)
		if err != nil {
			return []domain.Client{}, fmt.Errorf("failed to cast item to map[string]interface{}")
		}
		clients = append(clients, client)
	}
	return clients, nil
}
