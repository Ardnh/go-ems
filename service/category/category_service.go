package service

import (
	"context"

	"github.com/Ardnh/go-ems/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
