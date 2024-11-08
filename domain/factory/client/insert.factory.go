package factory

import (
	"study/domain/dto"
	"study/infra/repository/client"
)

type ClientFactory struct{}

func (factory ClientFactory) InsertFactory(dto dto.InsertDTO) client.InsertClient {
	return client.InsertClient{
		Name: dto.Name,
		Age:  dto.Age,
		Address: client.Address{
			City:    dto.Address.City,
			Zipcode: dto.Address.Zipcode,
		},
	}
}
