package registerevent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ss49919201/diary/user/domain/entity/registeruserevent"
	"github.com/ss49919201/diary/user/domain/entity/user"
	registerusereventRepo "github.com/ss49919201/diary/user/repository/registeruserevent"
)

var _ registerusereventRepo.Repository = (*Repository)(nil)

type Repository struct {
	db *dynamo.DB
}

func NewRepository() *Repository {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-west-2")})
	return &Repository{db}
}

func (r *Repository) Add(registeruserevent.RegisterUserEvent) (registeruserevent.RegisterUserEvent, error)
func (r *Repository) Exists(user.UserID) (bool, error)
func (r *Repository) Find(user.UserID) (registeruserevent.RegisterUserEvent, error)
