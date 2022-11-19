package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ardnh/go-ems/exception"
	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/web"
	service "github.com/Ardnh/go-ems/service/advertisement"
	"github.com/gorilla/mux"
)

type AdvertiseControllerImpl struct {
	Service service.AdvertisementService
}

func NewAdvertsementController(service service.AdvertisementService) AdvertisementController {
	return &AdvertiseControllerImpl{
		Service: service,
	}
}

func (controller *AdvertiseControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	userId, _ := strconv.Atoi(r.PostFormValue("user_id"))
	eventId, _ := strconv.Atoi(r.PostFormValue("event_id"))
	instagramUrl := r.PostFormValue("instagram_url")
	twitterUrl := r.PostFormValue("twitter_url")
	facebookUrl := r.PostFormValue("facebook_url")
	file, _, errParseImage := r.FormFile("banner")
	if errParseImage != nil {
		exception.InternalServerError(w, r, errParseImage)
	}

	url, errUpload := helper.ImageUpload(file, "CLOUDINARY_ADVERTISE_FOLDER_NAME")
	if errParseImage != nil {
		exception.InternalServerError(w, r, errUpload)
	}

	fmt.Println(url)

	advertiseCreateRequest := web.AdvertiseCreateRequest{
		UserId:       userId,
		EventId:      eventId,
		InstagramUrl: instagramUrl,
		TwitterUrl:   twitterUrl,
		FacebookUrl:  facebookUrl,
		BannerUrl:    url,
	}

	advertiseResponse := controller.Service.Create(r.Context(), advertiseCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   advertiseResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

// on progress
func (controller *AdvertiseControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
}

func (controller *AdvertiseControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	deleteRequest := web.AdvertiseDeleteRequest{}
	helper.ReadFromRequestBody(r, &deleteRequest)

	_, err := helper.DeleteImage(deleteRequest.PubId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	controller.Service.Delete(r.Context(), deleteRequest.Id, deleteRequest.UserId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AdvertiseControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	advertise := controller.Service.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   advertise,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AdvertiseControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	advertise := controller.Service.FindByUserId(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   advertise,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AdvertiseControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	advertise := controller.Service.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   advertise,
	}

	helper.WriteToResponseBody(w, webResponse)
}
