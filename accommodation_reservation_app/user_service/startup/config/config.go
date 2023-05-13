package config

import "os"

type Config struct {
	Port              string
	UserDBHost        string
	UserDBPort        string
	AccommodationHost string
	AccommodationPort string
	ReservationHost   string
	ReservationPort   string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:        os.Getenv("USER_DB_HOST"),
		UserDBPort:        os.Getenv("USER_DB_PORT"),
		AccommodationHost: os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort: os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		ReservationHost:   os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
