package monitoring

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/helper"
)

type MonitoringRepositoryImpl struct {
}

func NewMonitoringRepository() MonitoringRepository {
	return &MonitoringRepositoryImpl{}
}

func (repository *MonitoringRepositoryImpl) GetTotalUser(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(id) as total_user FROM user;"
	row := tx.QueryRowContext(ctx, SQL)

	var total int
	err := row.Scan(&total)
	helper.PanicIfError(err)

	return total
}

func (repository *MonitoringRepositoryImpl) GetTotalEvents(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(id) as total_events FROM events;"
	row := tx.QueryRowContext(ctx, SQL)

	var total int
	err := row.Scan(&total)
	helper.PanicIfError(err)

	return total
}

func (repository *MonitoringRepositoryImpl) GetTotalAdvertisement(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(id) as total_advertisement FROM advertisement;"
	row := tx.QueryRowContext(ctx, SQL)

	var total int
	err := row.Scan(&total)
	helper.PanicIfError(err)

	return total
}
