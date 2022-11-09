package user

import (
	"database/sql"

	"github.com/Ardnh/go-ems/repository/user"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Repository user.UserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewUserRepository(repository user.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}
