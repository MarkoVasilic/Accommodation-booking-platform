package main

import (
	"os"

	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/startup"
	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/startup/config"
)

func main() {
	if os.Getenv("RUN_ENV") != "production" {
		config.LoadEnvVariables()
	}
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
