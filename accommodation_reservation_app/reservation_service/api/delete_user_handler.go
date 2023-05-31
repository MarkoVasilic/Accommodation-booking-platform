package api

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	events "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/delete_user"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type DeleteUserCommandHandler struct {
	reservationService   *service.ReservationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
	accommodation_client accommodation_service.AccommodationServiceClient
	mongo_client         *mongo.Client
}

var reservations []models.Reservation

func NewDeleteUserCommandHandler(reservationService *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber, accommodation_client accommodation_service.AccommodationServiceClient, mongo_client *mongo.Client) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		reservationService:   reservationService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
		accommodation_client: accommodation_client,
		mongo_client:         mongo_client,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {
	reply := events.DeleteUserReply{User: command.User}
	id := command.User.Id
	token := command.User.Token
	switch command.Type {
	case events.DeleteReservations:
		_, err := handler.DeleteReservationsHost(id, token)
		if err != nil {
			reply.Type = events.ReservationsNotDeleted
			break
		}
		reply.Type = events.ReservationsDeleted
	case events.RollbackReservations:
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		reservationDocuments := make([]interface{}, len(reservations))
		for i, r := range reservations {
			reservationDocuments[i] = r
		}
		handler.reservationService.ReservationRepository.ReservationCollection.InsertMany(ctx, reservationDocuments)
		reply.Type = events.ReservationsRolledback
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

func (handler *DeleteUserCommandHandler) DeleteReservationsHost(id string, token string) (string, error) {
	allAccommodations, err := handler.accommodation_client.GetAllAccommodations(metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+token)),
		&accommodation_service.GetAllAccommodationsRequest{Id: id})
	if err != nil {
		return "Failed to get accommodations", err
	}

	var hostAccomodationsIds []string
	for _, accommodation := range allAccommodations.Accommodations {
		if accommodation.HostId == id {
			hostAccomodationsIds = append(hostAccomodationsIds, accommodation.Id)
		}
	}

	allAvailabilities, err := handler.accommodation_client.GetAllAvailabilities(metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+token)),
		&accommodation_service.GetAllAvailabilitiesRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return "Failed to get availabilities", err
	}

	var hostAvailabilitiesIds []primitive.ObjectID
	for _, availability := range allAvailabilities.Availabilities {
		for _, accommodationId := range hostAccomodationsIds {
			if availability.AccommodationID == accommodationId {
				id, err := primitive.ObjectIDFromHex(availability.Id)
				if err != nil {
					err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
					return "the provided id is not a valid ObjectID", err
				}
				hostAvailabilitiesIds = append(hostAvailabilitiesIds, id)
			}
		}
	}

	var hostReservations []models.Reservation
	for _, availabilityId := range hostAvailabilitiesIds {
		res, err := handler.reservationService.GetAll()
		if err != nil {
			return "something went wrong", err
		}

		for _, r := range res {
			if r.AvailabilityID == availabilityId {
				hostReservations = append(hostReservations, r)
			}
		}
	}
	reservations = hostReservations
	_, err = handler.reservationService.DeleteReservationsHost(hostReservations)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	return "", nil
}
