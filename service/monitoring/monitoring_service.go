package monitoring

import "context"

type MonitoringService interface {
	GetTotalEvents(ctx context.Context) int
	GetTotalAdvertisement(ctx context.Context) int
	GetTotalUser(ctx context.Context) int
}
