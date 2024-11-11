package domain

type AuthService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}

type HashService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, password string) error
}
