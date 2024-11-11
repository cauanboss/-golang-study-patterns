package auth

import (
	domain "study/domain/interfaces"
	"study/infra/repository/login"
)

type LoginUseCase struct {
	authService domain.AuthService
	hashService domain.HashService
	repository  login.RepositoryInterface
}

// TODO: FAZER DE VERDADE O FIND
func (uc *LoginUseCase) Login(username, password string) (string, error) {
	_, err := uc.repository.Find(username)
	if err != nil {
		return "", err
	}

	if err := uc.hashService.VerifyPassword("user.Password", password); err != nil {
		return "", err
	}

	return uc.authService.GenerateToken("user.ID")
}

// TODO: USAR O MAPPING/FACTORY
func (uc *LoginUseCase) insertOne(username, password string) (bool, error) {
	hashedPassword, err := uc.hashService.HashPassword(password)
	if err != nil {
		return false, err
	}

	dto := login.Insert{
		Username: username,
		Password: hashedPassword,
	}

	return uc.repository.InsertOne(dto)
}
