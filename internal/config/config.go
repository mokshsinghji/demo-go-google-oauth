package config

import (
	"github.com/joho/godotenv"
	"os"
)

func InitEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("Please make sure that an env file exists.")
	}
}

var (
	GOOGLE_CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	REDIRECT_URL         = "http://localhost:3000/login/google/callback"
)
