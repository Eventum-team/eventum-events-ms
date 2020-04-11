package service

import (
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
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
		u.Error(w, err,http.StatusInternalServerError)
		return
	}
	u.Respond(w, locations,http.StatusOK)
}


