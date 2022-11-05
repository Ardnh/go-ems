package web

type EventsCreateRequest struct {
	UserId                int    `json:"user_id" validate:"required"`
	CategoryId            int    `json:"category_id" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	Tagline               string `json:"tagline" validate:"required"`
	Description           string `json:"description" validate:"required"`
	Organizer             string `json:"organizer" validate:"required"`
	StartDate             string `json:"start_date" validate:"required datetime"`
	EndDate               string `json:"end_date" validate:"required datetime"`
	RegistrationStartDate string `json:"registration_start_date" validate:"required datetime"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required datetime"`
	RegistrationUrl       string `json:"registration_url" validate:"required"`
	Location              string `json:"location" validate:"required"`
	Capacity              int    `json:"capacity" validate:"required"`
	BannerUrl             string `json:"banner_url" validate:"required"`
	Visitor               int    `json:"visitor" validate:"required"`
	Status                string `json:"string" validate:"required"`
}

type EventsUpdateRequest struct {
	Id                    int    `json:"id" validate:"required"`
	UserId                int    `json:"user_id" validate:"required"`
	CategoryId            int    `json:"category_id" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	Tagline               string `json:"tagline" validate:"required"`
	Description           string `json:"description" validate:"required"`
	Organizer             string `json:"organizer" validate:"required"`
	StartDate             string `json:"start_date" validate:"required datetime"`
	EndDate               string `json:"end_date" validate:"required datetime"`
	RegistrationStartDate string `json:"registration_start_date" validate:"required datetime"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required datetime"`
	RegistrationUrl       string `json:"registration_url" validate:"required"`
	Location              string `json:"location" validate:"required"`
	Capacity              int    `json:"capacity" validate:"required"`
	BannerUrl             string `json:"banner_url" validate:"required"`
	Visitor               int    `json:"visitor" validate:"required"`
	Status                string `json:"string" validate:"required"`
}

type EventsResponse struct {
	Id                    int    `json:"id" validate:"required"`
	UserId                int    `json:"user_id" validate:"required"`
	CategoryId            int    `json:"category_id" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	Tagline               string `json:"tagline" validate:"required"`
	Description           string `json:"description" validate:"required"`
	Organizer             string `json:"organizer" validate:"required"`
	StartDate             string `json:"start_date" validate:"required datetime"`
	EndDate               string `json:"end_date" validate:"required datetime"`
	RegistrationStartDate string `json:"registration_start_date" validate:"required datetime"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required datetime"`
	RegistrationUrl       string `json:"registration_url" validate:"required"`
	Location              string `json:"location" validate:"required"`
	Capacity              int    `json:"capacity" validate:"required"`
	BannerUrl             string `json:"banner_url" validate:"required"`
	Visitor               int    `json:"visitor" validate:"required"`
	Status                string `json:"string" validate:"required"`
}
