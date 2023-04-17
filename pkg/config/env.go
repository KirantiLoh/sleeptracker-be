package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	dsn       string
	secretKey string
)

func LoadEnv() {
	godotenv.Load()
	dsn = os.Getenv("DATABASE_URL")
	secretKey = os.Getenv("SECRET_KEY")

	if dsn == "" || secretKey == "" {
		panic("Invalid .env")
	}
}
