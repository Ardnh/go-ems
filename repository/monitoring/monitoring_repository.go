package monitoring

import (
	"context"
	"database/sql"
)

type MonitoringRepository interface {
	GetTotalUser(ctx context.Context, tx *sql.Tx) int
	GetTotalEvents(ctx context.Context, tx *sql.Tx) int
	GetTotalAdvertisement(ctx context.Context, tx *sql.Tx) int
}
