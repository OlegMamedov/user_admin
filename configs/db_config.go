package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type DBsettings struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func LoadDBConfig() DBsettings {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := DBsettings{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),

	}

	return config
}
