package advertisement

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
	"github.com/Ardnh/go-ems/repository/advertisement"
	"github.com/go-playground/validator/v10"
)

type AdvertisementServiceImpl struct {
	Repository advertisement.AdvertisementRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewAdvertisementService(repository advertisement.AdvertisementRepository, db *sql.DB, validate *validator.Validate) AdvertisementService {
	return &AdvertisementServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *AdvertisementServiceImpl) Create(ctx context.Context, request web.AdvertiseCreateRequest) web.AdvertiseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	advertise := domain.Advertisement{
		UserId:       request.UserId,
		EventId:      request.EventId,
		InstagramUrl: request.InstagramUrl,
		TwitterUrl:   request.TwitterUrl,
		FacebookUrl:  request.FacebookUrl,
		BannerUrl:    request.BannerUrl,
	}

	advertisement := service.Repository.Save(ctx, tx, advertise)

	return helper.ToAdvertiseResponse(advertisement)
}

func (service *AdvertisementServiceImpl) Update(ctx context.Context, request web.AdvertiseUpdateRequest) web.AdvertiseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	advertise, err := service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	advertise.InstagramUrl = request.InstagramUrl
	advertise.TwitterUrl = request.TwitterUrl
	advertise.FacebookUrl = request.FacebookUrl
	advertise.BannerUrl = request.BannerUrl

	advertiseUpdate := service.Repository.Update(ctx, tx, advertise)
	return helper.ToAdvertiseResponse(advertiseUpdate)
}

func (service *AdvertisementServiceImpl) Delete(ctx context.Context, advertiseId int, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.Repository.Delete(ctx, tx, advertiseId, userId)
}

func (service *AdvertisementServiceImpl) FindById(ctx context.Context, advertiseId int) web.AdvertiseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	advertise, err := service.Repository.FindById(ctx, tx, advertiseId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdvertiseResponse(advertise)
}

func (service *AdvertisementServiceImpl) FindAll(ctx context.Context) []web.AdvertiseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	advertise := service.Repository.FindAll(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdvertiseResponses(advertise)
}

func (service *AdvertisementServiceImpl) FindByUserId(ctx context.Context, userId int) []web.AdvertiseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	advertise := service.Repository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdvertiseResponses(advertise)
}
