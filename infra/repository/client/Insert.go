package client

import "fmt"

func (client *IndexClient) InsertOne(dto InsertClient) (bool, error) {
	err := client.db.InsertOne(client.collectionName, dto)
	if err != nil {
		fmt.Println("Error InsertOne: ", err)
		return false, err
	}
	return true, nil
}
