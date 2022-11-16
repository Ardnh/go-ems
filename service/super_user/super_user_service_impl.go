package service

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
	superuser "github.com/Ardnh/go-ems/repository/super_user"
	"github.com/go-playground/validator/v10"
)

type SuperUserServiceImpl struct {
	Repository superuser.SuperUserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewSuperUserService(repository superuser.SuperUserRepository, db *sql.DB, validate *validator.Validate) SuperUserService {
	return &SuperUserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *SuperUserServiceImpl) Register(ctx context.Context, request web.SuperUserCreateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if err != nil {
		helper.PanicIfError(err)
	}

	user := domain.SuperUser{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Password:  request.Password,
	}

	service.Repository.Register(ctx, tx, user)
}

func (service *SuperUserServiceImpl) FindByUsername(ctx context.Context, request string) (web.SuperUserResponseByUserName, error) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByUsername(ctx, tx, request)
	if err != nil {
		helper.PanicIfError(err)
		exception.NewNotFoundError(err.Error())
	}

	return helper.ToSuperUserResponseByUsername(user), nil
}
