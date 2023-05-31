package delete_user

type UserDetails struct {
	Id    string
	Token string
}

type DeleteUserCommandType int8

const (
	DeleteReservations DeleteUserCommandType = iota
	RollbackReservations
	DeleteAccommodations
	RollbackAccommodations
	DeleteUser
	CancelDeletingUser
	UnknownCommand
)

type DeleteUserCommand struct {
	User UserDetails
	Type DeleteUserCommandType
}

type DeleteUserReplyType int8

const (
	ReservationsDeleted DeleteUserReplyType = iota
	ReservationsNotDeleted
	ReservationsRolledback
	AccommodationsDeleted
	AccommodationsNotDeleted
	AccommodationsRolledback
	UserDeleted
	UserNotDeleted
	UnknownReply
)

type DeleteUserReply struct {
	User UserDetails
	Type DeleteUserReplyType
}
