package api

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	events "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/delete_user"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserCommandHandler struct {
	accommodationService *service.AccommodationService
	availabilityService  *service.AvailabilityService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

var accommodations []models.Accommodation
var availabilities []models.Availability

func NewDeleteUserCommandHandler(accommodationService *service.AccommodationService, availabilityService *service.AvailabilityService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		accommodationService: accommodationService,
		availabilityService:  availabilityService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
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
		_, err := handler.DeleteAccommodationsHost(id, token)
		if err != nil {
			reply.Type = events.AccommodationsNotDeleted
			break
		}
		reply.Type = events.AccommodationsDeleted
	case events.RollbackAccommodations:
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		accommodationDocuments := make([]interface{}, len(accommodations))
		for i, r := range accommodations {
			accommodationDocuments[i] = r
		}
		availabilityDocuments := make([]interface{}, len(availabilities))
		for i, r := range availabilities {
			availabilityDocuments[i] = r
		}
		handler.accommodationService.AccommodationRepository.AccommodationCollection.InsertMany(ctx, accommodationDocuments)
		handler.availabilityService.AvailabilityRepository.AvailabilityCollection.InsertMany(ctx, availabilityDocuments)
		reply.Type = events.AccommodationsRolledback
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

func (handler *DeleteUserCommandHandler) DeleteAccommodationsHost(id string, token string) (string, error) {
	hostID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "invalid HostId", status.Errorf(codes.InvalidArgument, "invalid HostId")
	}
	allAccommodations, err := handler.accommodationService.GetAllAccommodations(hostID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return "the provided id is not a valid ObjectID", err
	}
	var hostAccomodations []models.Accommodation
	for _, accommodation := range allAccommodations {
		if accommodation.HostID == hostID {
			hostAccomodations = append(hostAccomodations, accommodation)
		}
	}
	allAvailabilities, err := handler.availabilityService.GetAllAvailabilities()
	if err != nil {
		return "the provided id is not a valid ObjectID", err
	}

	var hostAvailabilities []models.Availability
	for _, availability := range allAvailabilities {
		for _, accommodation := range hostAccomodations {
			if availability.AccommodationID == accommodation.ID {
				hostAvailabilities = append(hostAvailabilities, availability)
			}
		}
	}
	availabilities = hostAvailabilities
	_, err = handler.availabilityService.DeleteAvailabilitiesHost(hostAvailabilities)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	accommodations = hostAccomodations
	_, err = handler.accommodationService.DeleteAccommodationsHost(hostAccomodations)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	return "success", nil
}
