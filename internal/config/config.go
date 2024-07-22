package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func InitEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("Please make sure that an env file exists.")
	}

	GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectUrl = "http://localhost:3000/login/google/callback"

	fmt.Printf("Google Client ID: %s, Google Client Secret: %s\n", GoogleClientId, GoogleClientSecret)
}

var (
	GoogleClientId     = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectUrl        = "http://localhost:3000/login/google/callback"
)
