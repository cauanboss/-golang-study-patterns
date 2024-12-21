package factory

import (
	"encoding/json"
	"fmt"
	"study/domain"
)

func (factory ClientFactory) FindFactoryFromDB() {

}

func (factory ClientFactory) FindAllFactoryFromDB(data []interface{}) ([]domain.Client, error) {
	var clients []domain.Client

	for _, item := range data {
		// Supondo que `item` seja um mapa genérico
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to cast item to map[string]interface{}")
		}

		// Converte o mapa em JSON e depois para a struct
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
