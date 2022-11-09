package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
