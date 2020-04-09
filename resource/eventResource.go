package resource

import "ev-events-ms/service"

func HandleEventRequests()  {
	router := GetRouter()
	router.HandleFunc("/events", service.GetEvents).Methods("GET")
	router.HandleFunc("/events", service.CreateEvent).Methods("POST")
	router.HandleFunc("/events/{id}/", service.GetEventById).Methods("GET")
	router.HandleFunc("/events/{id}/", service.EditEvent).Methods("PUT")
	router.HandleFunc("/events/{id}/", service.DeleteEvent).Methods("DELETE")



}