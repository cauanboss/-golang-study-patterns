package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	URI      string
	Username string
	Password string
	DBName   string
}

type HTTPConfig struct {
	Port string
}

type AuthConfig struct {
	JWTSecret string
}

type Config struct {
	Database   DatabaseConfig
	HTTP       HTTPConfig
	AuthConfig AuthConfig
}

var Env Config

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao carregar o arquivo .env")
	}

	Env := Config{
		Database: DatabaseConfig{
			URI:      os.Getenv("DATABASE_MONGODB_URI"),
			Username: os.Getenv("DATABASE_MONGODB_USERNAME"),
			Password: os.Getenv("DATABASE_MONGODB_PASSWORD"),
			DBName:   os.Getenv("DATABASE_DB_NAME"),
		},
		HTTP: HTTPConfig{
			Port: os.Getenv("HTTP_PORT"),
		},
		AuthConfig: AuthConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
	}

	return Env
}
