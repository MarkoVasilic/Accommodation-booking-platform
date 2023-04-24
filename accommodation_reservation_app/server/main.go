package main

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/startup"
	cfg "github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/startup/config"
)

func main() {
	initializer.LoadEnvVariables()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
