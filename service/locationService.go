package service

import (
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"fmt"
	"net/http"
)

func GetLocations(w http.ResponseWriter, r *http.Request) () {
	events,err := repository.GetLocations()
	var locations []models.Location
	for _,ev := range events {
		loc := models.Location{
			Latitude: ev.Latitude,
			Longitude: ev.Longitude,
			EventId: ev.ID,
			EventName: ev.Name,
		}
		locations = append(locations,loc)
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, locations)
}


