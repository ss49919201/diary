package registeruserevent

import (
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
)

type Repository interface {
	Add(registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error)
	Exists(user.UserID) (bool, error)
	Find(user.UserID) (registeruserevent.RegisterUserEvent, error)
}
