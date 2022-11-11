package user

import (
	"context"

	"github.com/Ardnh/go-ems/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest)
	FindByUsername(ctx context.Context, request string) (web.UserResponseByUserName, error)
	Update(ctx context.Context, request web.UserUpdateRequest) error
}
