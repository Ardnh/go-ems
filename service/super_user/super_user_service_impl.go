package superuser

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/web"
	superuserRepository "github.com/Ardnh/go-ems/repository/super_user"
	"github.com/go-playground/validator/v10"
)

type SuperUserServiceImpl struct {
	Repository superuserRepository.SuperUserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewSuperUserService(repository superuserRepository.SuperUserRepository, db *sql.DB, validate *validator.Validate) SuperUserService {
	return &SuperUserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *SuperUserServiceImpl) Save(ctx context.Context, request web.SuperUserCreateCategory)
