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

func mapAccommodation(accommodation *models.Accommodation) *pb.Accommodation {

	accommodationPb := &pb.Accommodation{
		Id:         accommodation.ID.Hex(),
		HostId:     accommodation.HostID.Hex(),
		Name:       accommodation.Name,
		Location:   accommodation.Location,
		Wifi:       accommodation.Wifi,
		Kitchen:    accommodation.Kitchen,
		AC:         accommodation.AC,
		ParkingLot: accommodation.ParkingLot,
		MinGuests:  int32(accommodation.MinGuests),
		MaxGuests:  int32(accommodation.MaxGuests),
		Images:     make([]string, 0, len(accommodation.Images)),
		AutoAccept: accommodation.AutoAccept,
	}
	//proveriti ovo za images
	for _, image := range accommodation.Images {
		if image != nil {
			accommodationPb.Images = append(accommodationPb.Images, *image)
		}
	}

	return accommodationPb
}
