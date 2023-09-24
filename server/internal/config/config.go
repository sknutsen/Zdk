package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthDomain       string
	AuthAudience     string
	AuthClientId     string
	AuthClientSecret string
	AuthCallbackUrl  string
	DbType           string
	DbHost           string
	DbName           string
	DbUser           string
	DbPass           string
	DbPort           string
	Port             string
}

func NewConfig() *Config {
	env := os.Getenv("ENV")
	if env != "Railway" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}
	}

	config := &Config{
		AuthDomain:       os.Getenv("AUTH0_DOMAIN"),
		AuthAudience:     os.Getenv("AUTH0_AUDIENCE"),
		AuthClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		AuthClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		AuthCallbackUrl:  os.Getenv("AUTH0_CALLBACK_URL"),

		DbType: os.Getenv("DB_TYPE"),
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbPort: os.Getenv("DB_PORT"),

		Port: os.Getenv("PORT"),
	}

	os.Clearenv()

	return config
}
