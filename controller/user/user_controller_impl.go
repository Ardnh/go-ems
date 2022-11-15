package controller

import (
	"net/http"
	"strconv"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/web"
	service "github.com/Ardnh/go-ems/service/user"
	"github.com/gorilla/mux"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	userRequest := web.UserCreateRequest{}
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

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "INVALID USER ID",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	userUpdateRequest.Id = id

	errUpdate := controller.Service.Update(r.Context(), userUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   errUpdate,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userUpdateRequest,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := web.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userLogin)

}
