package web

type AdvertiseCreateRequest struct {
	UserId       int    `json:"user_id" validate:"required"`
	EventId      int    `json:"event_id" validate:"required"`
	InstagramUrl string `json:"instagram_url"`
	TwitterUrl   string `json:"twitter_url"`
	FacebookUrl  string `json:"facebook_url"`
	BannerUrl    string `json:"banner_url" validate:"required"`
}

type AdvertiseUpdateRequest struct {
	Id           int    `json:"id" validate:"required"`
	UserId       int    `json:"user_id" validate:"required"`
	EventId      int    `json:"event_id" validate:"required"`
	InstagramUrl string `json:"instagram_url"`
	TwitterUrl   string `json:"twitter_url"`
	FacebookUrl  string `json:"facebook_url"`
	BannerUrl    string `json:"banner_url" validate:"required"`
}

type AdvertiseDeleteRequest struct {
	Id     int
	UserId int
	PubId  string
}

type AdvertiseResponse struct {
	Id           int    `json:"id" validate:"required"`
	UserId       int    `json:"user_id" validate:"required"`
	EventId      int    `json:"event_id" validate:"required"`
	InstagramUrl string `json:"instagram_url"`
	TwitterUrl   string `json:"twitter_url"`
	FacebookUrl  string `json:"facebook_url"`
	BannerUrl    string `json:"banner_url" validate:"required"`
}
