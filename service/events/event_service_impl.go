package events

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
	"github.com/Ardnh/go-ems/repository/event"
	"github.com/go-playground/validator/v10"
)

type EventsServiceImpl struct {
	Repository event.EventRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewEventsServiceImpl(repository event.EventRepository, db *sql.DB, validate *validator.Validate) EventsService {
	return &EventsServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *EventsServiceImpl) Create(ctx context.Context, request web.EventsCreateRequest) web.EventsResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event := domain.Event{
		UserId:                request.UserId,
		CategoryId:            request.CategoryId,
		Name:                  request.Name,
		Tagline:               request.Tagline,
		Description:           request.Description,
		Organizer:             request.Organizer,
		StartDate:             request.StartDate,
		EndDate:               request.EndDate,
		RegistrationStartDate: request.RegistrationStartDate,
		RegistrationEndDate:   request.RegistrationEndDate,
		RegistrationUrl:       request.RegistrationUrl,
		Location:              request.Location,
		Capacity:              request.Capacity,
		BannerUrl:             request.BannerUrl,
		Visitor:               request.Visitor,
		Status:                request.Status,
	}

	event = service.Repository.Save(ctx, tx, event)
	return helper.ToEventResponse(event)
}

func (service *EventsServiceImpl) Update(ctx context.Context, request web.EventsUpdateRequest) web.EventsResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.Repository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	event.Id = request.Id
	event.UserId = request.UserId
	event.CategoryId = request.CategoryId
	event.Name = request.Name
	event.Tagline = request.Tagline
	event.Description = request.Description
	event.Organizer = request.Organizer
	event.StartDate = request.StartDate
	event.EndDate = request.EndDate
	event.RegistrationStartDate = request.RegistrationStartDate
	event.RegistrationEndDate = request.RegistrationEndDate
	event.RegistrationUrl = request.RegistrationUrl
	event.Location = request.Location
	event.Capacity = request.Capacity
	event.BannerUrl = request.BannerUrl
	event.Visitor = request.Visitor
	event.Status = request.Status

	event = service.Repository.Update(ctx, tx, event)
	return helper.ToEventResponse(event)
}

func (service *EventsServiceImpl) UpdateEventVisitor(ctx context.Context, eventId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.Repository.UpdateVisitor(ctx, tx, eventId)
}

func (service *EventsServiceImpl) Delete(ctx context.Context, eventId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.Repository.FindById(ctx, tx, eventId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.Repository.Delete(ctx, tx, event)
}

func (service *EventsServiceImpl) FindAll(ctx context.Context) []web.EventsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	events := service.Repository.FindAll(ctx, tx)

	return helper.ToEventsResponses(events)
}

func (service *EventsServiceImpl) FindById(ctx context.Context, eventId int) web.EventsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.Repository.FindById(ctx, tx, eventId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToEventResponse(event)
}

func (service *EventsServiceImpl) FindByCategoryId(ctx context.Context, categoryId int) []web.EventsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	events := service.Repository.FindByCategoryId(ctx, tx, categoryId)

	return helper.ToEventsResponses(events)
}

func (service *EventsServiceImpl) FindByUserId(ctx context.Context, userId int) []web.EventsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	events := service.Repository.FindByCategoryId(ctx, tx, userId)

	return helper.ToEventsResponses(events)
}
