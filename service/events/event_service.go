package events

import (
	"context"

	"github.com/Ardnh/go-ems/model/web"
)

type EventsService interface {
	Create(ctx context.Context, request web.EventsCreateRequest) web.EventsResponse
	Update(ctx context.Context, request web.EventsUpdateRequest) web.EventsResponse
	UpdateEventVisitor(ctx context.Context, eventId int)
	Delete(ctx context.Context, eventId int)
	FindAll(ctx context.Context) []web.EventsResponse
	FindById(ctx context.Context, eventId int) web.EventsResponse
	FindByCategoryId(ctx context.Context, categoryId int) []web.EventsResponse
	FindByUserId(ctx context.Context, userId int) []web.EventsResponse
}
