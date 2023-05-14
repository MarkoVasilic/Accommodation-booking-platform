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

func mapFindAvailability(findAvailability *models.FindAvailability) *pb.FindAvailability {
	startDate := timestamppb.New(findAvailability.StartDate)
	endDate := timestamppb.New(findAvailability.EndDate)
	var imageUrls []string
	for _, img := range findAvailability.Images {
		imageUrls = append(imageUrls, *img)
	}

	findAvailabilityPb := &pb.FindAvailability{
		AccommodationId: findAvailability.AccommodationId.Hex(),
		AvailabilityID:  findAvailability.AvailabilityID.Hex(),
		HostId:          findAvailability.HostID.Hex(),
		Name:            findAvailability.Name,
		Location:        findAvailability.Location,
		Wifi:            findAvailability.Wifi,
		Kitchen:         findAvailability.Kitchen,
		AC:              findAvailability.AC,
		ParkingLot:      findAvailability.ParkingLot,
		Images:          imageUrls,
		StartDate:       startDate,
		EndDate:         endDate,
		TotalPrice:      findAvailability.TotalPrice,
		SinglePrice:     findAvailability.SinglePrice,
		IsPricePerGuest: findAvailability.IsPricePerGuest,
	}

	return findAvailabilityPb
}

func mapAccommodation(accommodation *models.Accommodation) *pb.Accommodation {
	var imageUrls []string
	for _, img := range accommodation.Images {
		imageUrls = append(imageUrls, *img)
	}

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
		Images:     imageUrls, //dodati slike
		AutoAccept: accommodation.AutoAccept,
	}

	return accommodationPb
}
