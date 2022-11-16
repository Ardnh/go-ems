package controller

import (
	"net/http"

	service "github.com/Ardnh/go-ems/service/monitoring"
)

type MonitoringControllerImpl struct {
	Service service.MonitoringService
}

func NewMonitoringController(service service.MonitoringService) MonitoringController {
	return &MonitoringControllerImpl{
		Service: service,
	}
}

func (controller *MonitoringControllerImpl) GetTotalUser(w http.ResponseWriter, r *http.Request) {

}

func (controller *MonitoringControllerImpl) GetTotalEvents(w http.ResponseWriter, r *http.Request) {

}

func (controller *MonitoringControllerImpl) GetTotalAdvertise(w http.ResponseWriter, r *http.Request) {

}
