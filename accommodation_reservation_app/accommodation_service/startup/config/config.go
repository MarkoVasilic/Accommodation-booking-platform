package config

import "os"

type Config struct {
	Port                     string
	AccommodationDBHost      string
	AccommodationDBPort      string
	UserHost                 string
	UserPort                 string
	ReservationHost          string
	ReservationPort          string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:      os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:      os.Getenv("ACCOMMODATION_DB_PORT"),
		UserHost:                 os.Getenv("USER_SERVICE_HOST"),
		UserPort:                 os.Getenv("USER_SERVICE_PORT"),
		ReservationHost:          os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:          os.Getenv("RESERVATION_SERVICE_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
