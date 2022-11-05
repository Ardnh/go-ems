package helper

import (
	"github.com/Ardnh/go-ems/model/domain"
	"github.com/Ardnh/go-ems/model/web"
)

func ToSuperUserResponseByUsername(user domain.SuperUser) web.SuperUserResponseByUserName {
	return web.SuperUserResponseByUserName{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
	}
}

func ToEventResponse(event domain.Event) web.EventsResponse {
	return web.EventsResponse{
		Id:                    event.Id,
		UserId:                event.UserId,
		CategoryId:            event.CategoryId,
		Name:                  event.Name,
		Tagline:               event.Tagline,
		Description:           event.Description,
		Organizer:             event.Organizer,
		StartDate:             event.StartDate,
		EndDate:               event.EndDate,
		RegistrationStartDate: event.RegistrationStartDate,
		RegistrationEndDate:   event.RegistrationEndDate,
		RegistrationUrl:       event.RegistrationUrl,
		Location:              event.Location,
		Capacity:              event.Capacity,
		BannerUrl:             event.BannerUrl,
		Visitor:               event.Visitor,
		Status:                event.Status,
	}
}

func ToEventsResponses(events []domain.Event) []web.EventsResponse {
	var eventsResponses []web.EventsResponse
	for _, event := range events {
		eventsResponses = append(eventsResponses, ToEventResponse(event))
	}

	return eventsResponses
}
