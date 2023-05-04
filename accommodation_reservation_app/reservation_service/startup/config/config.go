package config

import "os"

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
	AccommodationHost string
	AccommodationPort string
	UserHost          string
	UserPort          string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		UserHost:          os.Getenv("USER_SERVICE_HOST"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
	}
}
