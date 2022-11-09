package advertisement

import (
	"context"

	"github.com/Ardnh/go-ems/model/web"
)

type AdvertisementService interface {
	Create(ctx context.Context, request web.AdvertiseCreateRequest) web.AdvertiseResponse
	Update(ctx context.Context, request web.AdvertiseUpdateRequest) web.AdvertiseResponse
	Delete(ctx context.Context, advertiseId int, userId int)
	FindById(ctx context.Context, advertiseId int) web.AdvertiseResponse
	FindAll(ctx context.Context) []web.AdvertiseResponse
	FindByUserId(ctx context.Context, userId int) []web.AdvertiseResponse
}
