package config

import "os"

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
	UserHost            string
	UserPort            string
	ReservationHost     string
	ReservationPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
		UserHost:            os.Getenv("USER_SERVICE_HOST"),
		UserPort:            os.Getenv("USER_SERVICE_PORT"),
		ReservationHost:     os.Getenv("USER_SERVICE_HOST"),
		ReservationPort:     os.Getenv("USER_SERVICE_PORT"),
	}
}
