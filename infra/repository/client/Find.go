package client

func (client *IndexClient) Find(id string) ([]interface{}, error) {
	if id == "" {
		all, err := client.db.FindAll(client.collectionName)
		if err != nil {
			return nil, err
		}
		return all, nil
	}
	one, err := client.db.FindOne(client.collectionName, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return []interface{}{one}, nil
}
