package service

import (
	"errors"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"net/http"
)


func GetEventsByStatus(w http.ResponseWriter, r *http.Request)  {
	status := r.URL.Query().Get("status")
	for _, opt := range models.StatusOptions {
		if opt == status {
			events,err := repository.GetEventsByStatus(opt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Error(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	u.Error(w, errors.New(status + " is not a valid status"))
}

func GetEventsByOwnerType(w http.ResponseWriter, r *http.Request)  {
	ownerType := r.URL.Query().Get("type")
	for _, opt := range models.OwnerTypeOptions {
		if opt == ownerType {
			events,err := repository.GetEventsByOwnerType(opt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Error(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	u.Error(w,errors.New(ownerType + " is not a valid owner type"))
}
//"2014-11-12T11:45:26Z"    -> Date Format
func GetEventsByRangeDate(w http.ResponseWriter, r *http.Request)  {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	valid1 := models.ValidateStringDate(start)
	valid2 := models.ValidateStringDate(end)

		if valid1 && valid2  {
			events,err := repository.GetEventsByRangeDate(start,end)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Error(w, err)
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	events := make([]*models.Event, 0)
	w.WriteHeader(http.StatusBadRequest)
	u.Respond(w, events)
}

func GetEventsByName(w http.ResponseWriter, r *http.Request)  {

	name := r.URL.Query().Get("name")
	events,err := repository.GetEventsByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, events)
	return
}


