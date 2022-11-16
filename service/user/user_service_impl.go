package service

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
	"github.com/Ardnh/go-ems/repository/user"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Repository user.UserRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewUserService(repository user.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		UserName:     request.UserName,
		Organization: request.Organization,
		Password:     request.Password,
	}

	service.Repository.Register(ctx, tx, user)
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, request string) (web.UserResponseByUserName, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByUsername(ctx, tx, request)
	if err != nil {
		helper.PanicIfError(err)
		exception.NewNotFoundError(err.Error())
	}

	return helper.ToUserResponseByUsername(user), nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) error {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByUsername(ctx, tx, request.UserName)
	if err != nil {
		helper.PanicIfError(err)
		exception.NewNotFoundError(err.Error())
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserName = request.UserName
	user.Organization = request.Organization

	err = service.Repository.Update(ctx, tx, user)
	return err
}
