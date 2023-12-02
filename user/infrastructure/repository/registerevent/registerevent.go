package registerevent

import (
	"github.com/guregu/dynamo"
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
	registerusereventRepo "github.com/ss49919201/diary/user/repository/registeruserevent"
	"github.com/uptrace/bun"
)

var _ registerusereventRepo.Repository = (*Repository)(nil)

type Repository struct {
	ddb *dynamo.DB
	db  *bun.DB
}

func NewRepository(ddb *dynamo.DB, db *bun.DB) *Repository {
	return &Repository{ddb, db}
}

func (r *Repository) Add(registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error)
func (r *Repository) Exists(user.UserID) (bool, error)
func (r *Repository) Find(user.UserID) (registeruserevent.RegisterUserEvent, error)
