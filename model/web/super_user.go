package web

type SuperUserCreateCategory struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SuperUserResponse struct {
	Id        string
	FirstName string
	LastName  string
	Username  string
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
