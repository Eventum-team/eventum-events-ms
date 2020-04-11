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
				u.Error(w, err,http.StatusInternalServerError)
				return
			}
			u.Respond(w, events,http.StatusOK)
			return
		}
	}
	u.Error(w, errors.New(status + " is  a valid status"),http.StatusBadRequest)
}

func GetEventsByOwnerType(w http.ResponseWriter, r *http.Request)  {
	ownerType := r.URL.Query().Get("type")
	for _, opt := range models.OwnerTypeOptions {
		if opt == ownerType {
			events,err := repository.GetEventsByOwnerType(opt)
			if err != nil {
				u.Error(w, err,http.StatusInternalServerError)
				return
			}
			u.Respond(w, events,http.StatusOK)
			return
		}
	}
	u.Error(w,errors.New(ownerType + " is not a valid owner type"),http.StatusBadRequest)
}
//"2014-11-12T11:45:26Z"    -> Date Format
func GetEventsByRangeDate(w http.ResponseWriter, r *http.Request)  {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	d1,valid1 := models.ValidateStringDate(start)
	d2,valid2 := models.ValidateStringDate(end)
		if valid1 && valid2  {
			errDates := models.ProperDates(d1,d2)
			if errDates != nil{
				u.Error(w, errDates,http.StatusBadRequest)
				return
			}
			events,err := repository.GetEventsByRangeDate(start,end)
			if err != nil {
				u.Error(w, err,http.StatusInternalServerError)
			}
			u.Respond(w, events,http.StatusOK)
			return
		}
	u.Respond(w, u.Message("Invalid date format ",http.StatusBadRequest),http.StatusBadRequest)
}

func GetEventsByName(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")
	events,err := repository.GetEventsByName(name)
	if err != nil {
		u.Error(w, err,http.StatusInternalServerError)
		return
	}
	u.Respond(w, events,http.StatusOK)
}


