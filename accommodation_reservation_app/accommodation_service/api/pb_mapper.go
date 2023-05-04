package api

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapAvailability(availability *models.Availability) *pb.Availability {
	startDate := timestamppb.New(availability.StartDate)
	endDate := timestamppb.New(availability.EndDate)

	availabilityPb := &pb.Availability{
		Id:              availability.ID.Hex(),
		AccommodationID: availability.AccommodationID.Hex(),
		StartDate:       startDate,
		EndDate:         endDate,
		Price:           availability.Price,
		IsPricePerGuest: availability.IsPricePerGuest,
	}
	return availabilityPb
}
