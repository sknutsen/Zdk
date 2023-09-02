package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthDomain   string
	AuthAudience string
	DbType       string
	DbHost       string
	DbName       string
	DbUser       string
	DbPass       string
	DbPort       string
}

func NewConfig() *Config {
	config := &Config{}

	config.LoadEnv()

	return config
}

func (config *Config) LoadEnv() {
	env := os.Getenv("ENV")
	if env != "Railway" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}
	}

	config.AuthDomain = os.Getenv("AUTH0_DOMAIN")
	config.AuthAudience = os.Getenv("AUTH0_AUDIENCE")

	config.DbType = os.Getenv("DB_TYPE")
	config.DbHost = os.Getenv("DB_HOST")
	config.DbName = os.Getenv("DB_NAME")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPass = os.Getenv("DB_PASS")
	config.DbPort = os.Getenv("DB_PORT")

	os.Clearenv()
}
