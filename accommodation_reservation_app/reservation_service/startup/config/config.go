package config

import "os"

type Config struct {
	Port                     string
	ReservationDBHost        string
	ReservationDBPort        string
	AccommodationHost        string
	AccommodationPort        string
	UserHost                 string
	UserPort                 string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost:        os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort:        os.Getenv("RESERVATION_DB_PORT"),
		AccommodationHost:        os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:        os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		UserHost:                 os.Getenv("USER_SERVICE_HOST"),
		UserPort:                 os.Getenv("USER_SERVICE_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
