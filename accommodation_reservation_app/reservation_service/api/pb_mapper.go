package api

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapReservation(reservation *models.Reservation) *pb.Reservation {
	startDate := timestamppb.New(reservation.StartDate)
	endDate := timestamppb.New(reservation.EndDate)

	resevationPb := &pb.Reservation{
		Id:             reservation.ID.Hex(),
		AvailabilityID: reservation.AvailabilityID.Hex(),
		GuestId:        reservation.GuestID.Hex(),
		StartDate:      startDate,
		EndDate:        endDate,
		NumGuests:      int32(reservation.NumGuests),
		IsAccepted:     reservation.IsAccepted,
		IsCanceled:     reservation.IsCanceled,
		IsDeleted:      reservation.IsDeleted,
	}
	return resevationPb
}

func mapFindReservation(findReservation *models.FindReservation) *pb.FindReservation {
	startDate := timestamppb.New(findReservation.StartDate)
	endDate := timestamppb.New(findReservation.EndDate)

	resevationPb := &pb.FindReservation{
		ReservationId:    findReservation.ReservationId.Hex(),
		GuestId:          findReservation.GuestID.Hex(),
		Name:             findReservation.Name,
		Location:         findReservation.Location,
		StartDate:        startDate,
		EndDate:          endDate,
		NumOfCancelation: int32(findReservation.NumOfCancelation), //da li na 0 ovde ili negde drugo?
		IsAccepted:       findReservation.IsAccepted,
		IsCanceled:       findReservation.IsCanceled,
	}
	return resevationPb
}
