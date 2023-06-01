package main

import (
	"os"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/startup"
	cfg "github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/startup/config"
)

func main() {
	if os.Getenv("RUN_ENV") != "production" {
		initializer.LoadEnvVariables()
	}
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
