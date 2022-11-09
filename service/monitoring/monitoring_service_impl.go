package monitoring

import (
	"database/sql"

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
