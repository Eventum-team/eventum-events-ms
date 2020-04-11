package resource

import "ev-events-ms/service"

func HandleEventFilterRequests()  {
	router := GetRouter()

	router.HandleFunc("/events/filter/locations", service.GetLocations).Methods("GET")
	router.HandleFunc("/events/filter/status", service.GetEventsByStatus).Methods("GET")
	router.HandleFunc("/events/filter/ownerType", service.GetEventsByOwnerType).Methods("GET")
	router.HandleFunc("/events/filter/rangeDate", service.GetEventsByRangeDate).Methods("GET")
	router.HandleFunc("/events/filter/name", service.GetEventsByName).Methods("GET")


}