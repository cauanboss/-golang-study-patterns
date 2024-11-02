package client

import "fmt"

type Address struct {
	City    string
	Zipcode string
}

// TODO: Change the struct name to InsertClient
type InsertClient struct {
	Name    string
	Age     int
	Address Address
}

func (client *IndexClient) InsertOne(dto InsertClient) (bool, error) {
	err := client.db.InsertOne(client.collectionName, dto)
	if err != nil {
		fmt.Println("Error InsertOne: ", err)
		return false, err
	}
	return true, nil
}
