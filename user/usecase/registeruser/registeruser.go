package registeruser

import (
	"errors"
	"fmt"
	"time"

	"github.com/ss49919201/diary/id"
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
)

type Input struct{}

type Output struct {
	UserID       string
	RegisteredAt time.Time
}

type Usecase interface {
	Run(Input) (Output, error)
}

type ExistsRegisterUserEvent func(user.UserID) (bool, error)

type AddRegisterUserEvent func(registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error)

type usecase struct {
	existsRegisterUserEvent ExistsRegisterUserEvent
	addRegisterUserEvent    AddRegisterUserEvent
	generateID              id.Generate
}

func (u *usecase) Run(in Input) (Output, error) {
	w := new(workflow).
		newUser(u.generateID).
		registerUser(u.existsRegisterUserEvent, u.addRegisterUserEvent).
		toOutput()
	return w.output, w.err
}

type workflow struct {
	err    error
	u      user.User
	re     registeruserevent.RegisterUserEvent
	output Output
}

func (w *workflow) addError(err error) *workflow {
	if err == nil {
		return w
	}

	if w.err == nil {
		w.err = err
	} else {
		w.err = fmt.Errorf("%v; %w", w.err, err)
	}

	return w
}

func (w *workflow) newUser(generateID id.Generate) *workflow {
	w.u = user.NewUser(user.NewUserID(generateID()))
	return w
}

func (w *workflow) registerUser(
	existsRegisterUserEvent ExistsRegisterUserEvent,
	addRegisterUserEvent AddRegisterUserEvent,
) *workflow {
	ok, err := existsRegisterUserEvent(w.u.UserID())
	if err != nil {
		w.addError(err)
		return w
	}
	if ok {
		w.addError(errors.New("already exists"))
		return w
	}

	event := registeruserevent.NewRegisterUserEvent(w.u, time.Now())
	event, err = addRegisterUserEvent(event)
	if err != nil {
		w.addError(err)
		return w
	}
	w.re = event

	return w
}

func (w *workflow) toOutput() *workflow {
	w.output = Output{
		UserID: func() string {
			uid := w.re.UserID()
			return uid.Value()
		}(),
		RegisteredAt: w.re.RegisteredAt(),
	}
	return w
}
