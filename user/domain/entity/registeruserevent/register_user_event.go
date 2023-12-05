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

func (r *RegisterUserEvent) UserID() user.UserID {
	return r.registeredUser.UserID()
}

func (r *RegisterUserEvent) RegisteredAt() time.Time {
	return r.registeredAt
}
