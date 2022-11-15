package web

type UserCreateRequest struct {
	FirstName    string `json:"firstname" validate:"required"`
	LastName     string `json:"lastname" validate:"required"`
	UserName     string `json:"username" validate:"required"`
	Organization string `json:"organization" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Id           int    `json:"id" validate:"required"`
	FirstName    string `json:"firstname" validate:"required"`
	LastName     string `json:"lastname" validate:"required"`
	UserName     string `json:"username" validate:"required"`
	Organization string `json:"organization" validate:"required"`
}

type UserResponseByUserName struct {
	Id           int
	FirstName    string
	LastName     string
	UserName     string
	Organization string
	Password     string
}

type ResponseUserWithToken struct {
	Id           int    `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	UserName     string `json:"username"`
	Organization string `json:"organization"`
	Token        string `json:"token"`
}

type UserResponseWithToken struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   ResponseWithToken `json:"data"`
}

type UserLoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
