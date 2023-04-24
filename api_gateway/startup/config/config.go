package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	AccommodationHost string
	AccommodationPort string
}

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
	}
}
