package registeruser

import (
	"errors"
	"time"

	"github.com/ss49919201/diary/id"
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
)

type Input struct{}

type Output struct{}

type Usecase interface {
	Run(Input) (Output, error)
}

type ExistsRegisterUserEvent = func(user.UserID) (bool, error)

type AddRegisterUserEvent = func(registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error)

type usecase struct {
	existsRegisterUserEvent ExistsRegisterUserEvent
	addRegisterUserEvent    AddRegisterUserEvent
	generateID              id.Generate
}

func (u *usecase) Run(in Input) (Output, error) {
	out := Output{}
	newUser := user.NewUser(user.NewUserID(u.generateID()))
	ok, err := u.existsRegisterUserEvent(newUser.UserID())
	if err != nil {
		return out, err
	}
	if ok {
		return out, errors.New("already exists")
	}
	event := registeruserevent.NewRegisterUserEvent(newUser, time.Now())
	_, err = u.addRegisterUserEvent(event)
	if err != nil {
		return out, err
	}
	return out, nil
}
