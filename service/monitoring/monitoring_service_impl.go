package monitoring

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/repository/monitoring"
	"github.com/go-playground/validator/v10"
)

type MonitoringServiceImpl struct {
	Repository monitoring.MonitoringRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewMonitoringService(repository monitoring.MonitoringRepository, db *sql.DB, validate *validator.Validate) MonitoringService {
	return &MonitoringServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *MonitoringServiceImpl) GetTotalUser(ctx context.Context) int {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	totalUser := service.Repository.GetTotalUser(ctx, tx)

	return totalUser
}

func (service *MonitoringServiceImpl) GetTotalEvents(ctx context.Context) int {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	totalEvents := service.Repository.GetTotalEvents(ctx, tx)

	return totalEvents
}

func (service *MonitoringServiceImpl) GetTotalAdvertisement(ctx context.Context) int {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	totalAdvertise := service.Repository.GetTotalAdvertisement(ctx, tx)

	return totalAdvertise
}
