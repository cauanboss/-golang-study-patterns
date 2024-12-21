package client

import "fmt"

func (client *IndexClient) InsertOne(data Client) (bool, error) {
	err := client.db.InsertOne(client.collectionName, data)
	if err != nil {
		fmt.Println("Error InsertOne: ", err)
		return false, err
	}
	return true, nil
}
