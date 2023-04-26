package main

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/startup"
	cfg "github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/startup/config"
)

func main() {
	initializer.LoadEnvVariables()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
