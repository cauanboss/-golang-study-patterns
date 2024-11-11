package login

func (login *IndexLogin) InsertOne(dto Insert) (bool, error) {
	err := login.db.InsertOne(login.collectionName, dto)
	if err != nil {
		return false, err
	}
	return true, nil
}
