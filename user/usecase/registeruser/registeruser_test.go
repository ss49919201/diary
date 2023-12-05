package registeruser

import (
	"reflect"
	"testing"
	"time"

	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
)

func Test_usecase_Run(t *testing.T) {
	u := &usecase{
		existsRegisterUserEvent: func(ui user.UserID) (bool, error) {
			return false, nil
		},
		addRegisterUserEvent: func(rue registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error) {
			return registeruserevent.NewRegisterUserEvent(
				user.NewUser(user.NewUserID("dummy")),
				time.Time{},
			), nil
		},
		generateID: func() string {
			return "dummy"
		},
	}

	out, err := u.Run(Input{})
	if err != nil {
		t.Errorf("usecase.Run() expect no error, actual %v", err)
		return
	}

	expect := Output{
		UserID: "dummy",
	}
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("usecase.Run() want %v, actual %v", expect, out)
	}
}
