package service

import (
	events "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/delete_user"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserOrchestrator, error) {
	o := &DeleteUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeleteUserOrchestrator) Start(id primitive.ObjectID, ClientToken string) error {
	event := &events.DeleteUserCommand{
		Type: events.DeleteReservations,
		User: events.UserDetails{
			Id:    id.Hex(),
			Token: ClientToken,
		},
	}
	return o.commandPublisher.Publish(event)
}

func (o *DeleteUserOrchestrator) handle(reply *events.DeleteUserReply) {
	command := events.DeleteUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteUserOrchestrator) nextCommandType(reply events.DeleteUserReplyType) events.DeleteUserCommandType {
	switch reply {
	case events.ReservationsDeleted:
		return events.DeleteAccommodations
	case events.ReservationsNotDeleted:
		return events.CancelDeletingUser
	case events.ReservationsRolledback:
		return events.CancelDeletingUser
	case events.AccommodationsDeleted:
		return events.DeleteUser
	case events.AccommodationsNotDeleted:
		return events.RollbackReservations
	case events.AccommodationsRolledback:
		return events.CancelDeletingUser
	case events.UserNotDeleted:
		return events.RollbackAccommodations
	default:
		return events.UnknownCommand
	}
}
