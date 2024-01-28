package config

import (
	"fmt"
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
	DbUpdate         bool
	Port             string
	Host             string
}

func NewConfig() *Config {
	env := os.Getenv("ENV")
	if env != "Railway" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}
	}

	port := os.Getenv("PORT")

	var callbackUrl string

	host := os.Getenv("HOST")

	if host != "" {
		callbackUrl = fmt.Sprintf("http://%s/%s", os.Getenv("HOST"), os.Getenv("AUTH0_CALLBACK_URL"))
	} else {
		var hostName string

		if port == "" {
			port = "8080"
			hostName = "127.0.0.1"
		} else {
			hostName = "0.0.0.0"
		}

		callbackUrl = fmt.Sprintf("http://%s:%s/%s", hostName, port, os.Getenv("AUTH0_CALLBACK_URL"))
	}

	config := &Config{
		AuthDomain:       os.Getenv("AUTH0_DOMAIN"),
		AuthAudience:     os.Getenv("AUTH0_AUDIENCE"),
		AuthClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		AuthClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		AuthCallbackUrl:  callbackUrl,

		DbType:   os.Getenv("DB_TYPE"),
		DbHost:   os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		DbUser:   os.Getenv("DB_USER"),
		DbPass:   os.Getenv("DB_PASS"),
		DbPort:   os.Getenv("DB_PORT"),
		DbUpdate: os.Getenv("DB_UPDATE") != "false",

		Host: host,

		Port: port,
	}

	os.Clearenv()

	return config
}
