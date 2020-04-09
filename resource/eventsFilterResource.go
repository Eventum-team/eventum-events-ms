package resource

import "ev-events-ms/service"

func HandleEventFilterRequests()  {
	router := GetRouter()

	router.HandleFunc("/events/locations", service.GetLocations).Methods("GET")
	router.HandleFunc("/events/status", service.GetEventsByStatus).Methods("POST")
	router.HandleFunc("/events/ownerType", service.GetEventsByOwnerType).Methods("POST")
	router.HandleFunc("/events/upcoming", service.GetUpcomingEvents).Methods("POST")
	router.HandleFunc("/events/date", service.GetEventsByDate).Methods("POST")


}