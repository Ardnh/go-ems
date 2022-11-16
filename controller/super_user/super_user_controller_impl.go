package controller

import (
	"fmt"
	"net/http"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/web"
	service "github.com/Ardnh/go-ems/service/super_user"
)

type UserControllerImpl struct {
	Service service.SuperUserService
}

func NewSuperUserController(service service.SuperUserService) SuperUserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	userRequest := web.SuperUserCreateRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	user, _ := controller.Service.FindByUsername(r.Context(), userRequest.UserName)

	// can improve
	if user.UserName == userRequest.UserName {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "USER ALREADY EXIST",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	hashedPassword, err := helper.GenerateHashPassword(userRequest.Password)
	if err != nil {
		exception.InternalServerError(w, r, err)
	}

	userRequest.Password = hashedPassword

	controller.Service.Register(r.Context(), userRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := web.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userLogin)

	user, err := controller.Service.FindByUsername(r.Context(), userLogin.UserName)
	if err != nil {
		exception.NewNotFoundError(err.Error())
	}

	if helper.CheckPassword(userLogin.Password, user.Password) {
		id := fmt.Sprintf("%d", user.Id)
		token, err := helper.GenerateJWTKey(id)
		if err != nil {
			exception.InternalServerError(w, r, err)
		}

		webResponse := web.UserResponseWithToken{
			Code:   200,
			Status: "OK",
			Data: web.ResponseWithToken{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				UserName:  user.UserName,
				Token:     token,
			},
		}

		helper.WriteToResponseBody(w, webResponse)
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Wrong Password",
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
