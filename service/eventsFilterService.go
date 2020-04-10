package service

import (
	"encoding/json"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"net/http"
	"strings"
	"time"
)

type Message struct {
	OwnerType string
	StatusOption string
	Date1 time.Time
	Date2 time.Time
	EventName string
}


func GetEventsByStatus(w http.ResponseWriter, r *http.Request)  {
	 msg := &Message{}
	 decodeBody(w,r,msg)
	for _, opt := range models.StatusOptions {
		if opt == msg.StatusOption {
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
	events := make([]*models.Event, 0)
	w.WriteHeader(http.StatusBadRequest)
	u.Respond(w, events)
}

func GetEventsByOwnerType(w http.ResponseWriter, r *http.Request)  {
	msg := &Message{}
	decodeBody(w,r,msg)
	for _, opt := range models.OwnerTypeOptions {
		if opt == msg.OwnerType {
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
	events := make([]*models.Event, 0)
	w.WriteHeader(http.StatusBadRequest)
	u.Respond(w, events)
}
//"2014-11-12T11:45:26Z"    -> Date Format
func GetEventsByRangeDate(w http.ResponseWriter, r *http.Request)  {
	msg := &Message{}
	decodeBody(w,r,msg)
	valid1 := models.ValidateDate(msg.Date1)
	valid2 := models.ValidateDate(msg.Date2)

		if valid1 && valid2  {
			events,err := repository.GetEventsByRangeDate(msg.Date1,msg.Date2)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Error(w, err)
			}
			w.WriteHeader(http.StatusOK)
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	events := make([]*models.Event, 0)
	w.WriteHeader(http.StatusBadRequest)
	u.Respond(w, events)
}

func GetEventsByName(w http.ResponseWriter, r *http.Request)  {
	msg := &Message{}
	decodeBody(w,r,msg)
	events,err := repository.GetEventsByName(msg.EventName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, events)
	return
}

func decodeBody(w http.ResponseWriter,r *http.Request,msg *Message)  {
	if r.Body == nil {
		//http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {

		//http.Error(w, err.Error(), 400)
		return
	}
	msg.StatusOption = strings.ToLower(msg.StatusOption)
	msg.OwnerType = strings.ToLower(msg.OwnerType)
}



