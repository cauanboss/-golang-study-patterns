package client

import "fmt"

func (client *IndexClient) Find(id string) (interface{}, error) {

	if id == "" {
		all, err := client.db.FindAll(client.collectionName)

		fmt.Println(all, err)
		if err != nil {
			return nil, err
		}

		return all, nil

	}

	one, err := client.db.FindOne(client.collectionName, nil)
	if err != nil {
		return nil, err
	}

	return one, nil
}
