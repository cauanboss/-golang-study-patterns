package client

func (client *IndexClient) Find(id string) ([]interface{}, error) {
	if id == "" {
		// Busca todos os documentos quando o ID não é fornecido
		all, err := client.db.FindAll(client.collectionName)
		if err != nil {
			return nil, err
		}

		return all, nil
	}

	// Busca um único documento quando o ID é fornecido
	one, err := client.db.FindOne(client.collectionName, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	// Retorna o resultado como um slice contendo o único item encontrado
	return []interface{}{one}, nil
}
