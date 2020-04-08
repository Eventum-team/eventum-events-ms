package resource

import "ev-events-ms/service"

func HandleLocationRequests()  {
	router := GetRouter()

	router.HandleFunc("/location", service.GetLocations).Methods("GET")
	router.HandleFunc("/location/{id}", service.GetLocationByEventId).Methods(" GET")
	router.HandleFunc("/location", service.CreateLocation).Methods("POST")
	router.HandleFunc("/location/{id}", service.EditLocationByEventId).Methods(" PUT")
	router.HandleFunc("/location/{id}", service.DeleteLocationByEventId).Methods(" DELETE")
}