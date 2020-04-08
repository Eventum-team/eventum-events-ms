package resource

import "ev-events-ms/service"

func HandleLocationRequests()  {
	router := GetRouter()
	router.HandleFunc("/locations", service.GetLocations).Methods("GET")
	router.HandleFunc("/locations/{id}", service.GetLocationById).Methods(" GET")
	router.HandleFunc("/locations", service.CreateLocation).Methods("POST")
	router.HandleFunc("/locations/{id}", service.EditLocation).Methods(" PUT")
	router.HandleFunc("/locations/{id}", service.DeleteLocation).Methods(" DELETE")
}