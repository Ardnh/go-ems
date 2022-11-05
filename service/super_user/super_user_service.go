package superuser

import (
	"context"

	"github.com/Ardnh/go-ems/model/web"
)

type SuperUserService interface {
	Register(ctx context.Context, request web.SuperUserCreateRequest)
	FindByUsername(ctx context.Context, request string) (web.SuperUserResponseByUserName, error)
}
