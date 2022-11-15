package controller

import (
	"net/http"
	"strconv"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/web"
	service "github.com/Ardnh/go-ems/service/category"
	"github.com/gorilla/mux"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: service,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.Service.Create(r.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "INVALID CATEGORY ID",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	categoryUpdateRequest.Id = categoryId
	controller.Service.Update(r.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryUpdateRequest,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])

	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "INVALID CATEGORY ID",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	controller.Service.Delete(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])

	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "INVALID CATEGORY ID",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	category := controller.Service.FindById(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	category := controller.Service.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(w, webResponse)
}
