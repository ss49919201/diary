package registeruser

import (
	"errors"
	"time"

	"github.com/ss49919201/diary/id"
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
	registerusereventRepo "github.com/ss49919201/diary/user/repository/registeruserevent"
)

type Input struct{}

type Output struct{}

type Usecase interface {
	Run(Input) (Output, error)
}

type usecase struct {
	registerusereventRepo.Repository
	id.Generate
}

func (u *usecase) Run(in Input) (Output, error) {
	out := Output{}
	newUser := user.NewUser(user.NewUserID(u.Generate()))
	ok, err := u.Exists(newUser.UserID())
	if err != nil {
		return out, err
	}
	if ok {
		return out, errors.New("already exists")
	}
	event := registeruserevent.NewRegisterUserEvent(newUser, time.Now())
	_, err = u.Add(event)
	if err != nil {
		return out, err
	}
	return out, nil
}
