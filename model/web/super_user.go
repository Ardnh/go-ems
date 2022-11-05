package web

type SuperUserCreateRequest struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type SuperUserResponseByUserName struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

type ResponseWithToken struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Token     string `json:"token"`
}

type SuperUserResponseWithToken struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   ResponseWithToken `json:"data"`
}
