package registeruserevent

import (
	"time"

	"github.com/ss49919201/diary/user/domain/entity/user"
)

type RegisterUserEvent struct {
	registeredUser user.User
	registeredAt   time.Time
}

func NewRegisterUserEvent(registeredUser user.User, registeredAt time.Time) RegisterUserEvent {
	return RegisterUserEvent{
		registeredUser: registeredUser,
		registeredAt:   registeredAt,
	}
}
