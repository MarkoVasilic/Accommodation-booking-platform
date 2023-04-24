package main

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/startup"
	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/startup/config"
)

func main() {
	config.LoadEnvVariables()
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
