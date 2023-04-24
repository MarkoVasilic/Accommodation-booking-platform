module github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway

go 1.20

replace github.com/MarkoVasilic/Accommodation-booking-platform/common => ../common

require (
	github.com/MarkoVasilic/Accommodation-booking-platform/common v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
