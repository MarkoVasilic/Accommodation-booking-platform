package api

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	events "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/delete_user"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUserCommandHandler struct {
	userService          *service.UserService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
	accommodation_client accommodation_service.AccommodationServiceClient
	reservation_client   reservation_service.ReservationServiceClient
}

func NewDeleteUserCommandHandler(userService *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber, accommodation_client accommodation_service.AccommodationServiceClient, reservation_client reservation_service.ReservationServiceClient) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		userService:          userService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
		accommodation_client: accommodation_client,
		reservation_client:   reservation_client,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	reply := events.DeleteUserReply{User: command.User}

	switch command.Type {
	case events.DeleteUser:
		result, err := handler.userService.UserRepository.DeleteUser(id)
		if err != nil {
			reply.Type = events.UserNotDeleted
			return
		}

		if result.DeletedCount == 0 {
			reply.Type = events.UserNotDeleted
			return
		}
		reply.Type = events.UserDeleted
		return
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
