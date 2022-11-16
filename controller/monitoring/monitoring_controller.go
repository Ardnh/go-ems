package controller

import "net/http"

type MonitoringController interface {
	GetTotalUser(w http.ResponseWriter, r *http.Request)
	GetTotalEvents(w http.ResponseWriter, r *http.Request)
	GetTotalAdvertise(w http.ResponseWriter, r *http.Request)
}
