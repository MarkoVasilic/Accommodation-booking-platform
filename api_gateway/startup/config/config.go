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
	UserHost          string
	UserPort          string
	ReservationHost   string
	ReservationPort   string
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
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		ReservationHost:   os.Getenv("USER_SERVICE_HOST"),
		ReservationPort:   os.Getenv("USER_SERVICE_PORT"),
	}
}
