package login

import "fmt"

func (login *IndexLogin) Find(id string) (interface{}, error) {

	if id == "" {
		all, err := login.db.FindAll(login.collectionName)

		fmt.Println(all, err)
		if err != nil {
			return nil, err
		}

		return all, nil

	}

	one, err := login.db.FindOne(login.collectionName, nil)
	if err != nil {
		return nil, err
	}

	return one, nil
}
