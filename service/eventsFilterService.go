package service

import (
	"encoding/json"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	FilterType string
	Params string
}

var statusOptions = []string{
 "upcoming",
 "cancelled",
 "ended",
}
var ownerTypeOptions = []string{
	"group",
	"user",
}
func GetEventsByStatus(w http.ResponseWriter, r *http.Request)  {
	 msg := &Message{}
	 decodeBody(r,msg)
	 if msg.FilterType != "status"{
		 u.Respond(w, u.Message(false, "Wrong filter type"))
		 return
	 }
	for _, opt := range statusOptions {
		if opt == msg.Params {
			events,err := repository.GetEventsByStatus(opt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Respond(w, u.Message(false, "Database Connection error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	u.Respond(w, u.Message(false, msg.Params + " is not a status"))
}

func GetEventsByOwnerType(w http.ResponseWriter, r *http.Request)  {
	msg := &Message{}
	decodeBody(r,msg)
	if msg.FilterType != "ownertype"{
		u.Respond(w, u.Message(false, "Wrong filter type"))
		return
	}
	for _, opt := range ownerTypeOptions {
		if opt == msg.Params {
			events,err := repository.GetEventsByOwnerType(opt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Respond(w, u.Message(false, "Database Connection error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	u.Respond(w, u.Message(false, msg.Params + " is not a owner type"))
}

func GetUpcomingEvents(w http.ResponseWriter, r *http.Request)  {
	msg := &Message{}
	decodeBody(r,msg)
	if msg.FilterType != "upcoming"{
		u.Respond(w, u.Message(false, "Wrong filter type"))
		return
	}
	v := true
	//validateDate()

		if v  {
			events,err := repository.GetUpcomingEvents(msg.Params)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				u.Respond(w, u.Message(false, "Database Connection error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			u.Respond(w, events)
			return
		}
	w.WriteHeader(http.StatusNotFound)
	u.Respond(w, u.Message(false, msg.Params + " is not valid date"))
}

//"2014-11-12T11:45:26Z"    -> Date Format
func GetEventsByDate(w http.ResponseWriter, r *http.Request)  {
	events,err := repository.GetEvents()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, events)
}

func decodeBody(r *http.Request,msg *Message)  {

	if r.Body == nil {
		//http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {

		//http.Error(w, err.Error(), 400)
		return
	}
	msg.Params = strings.ToLower(msg.Params)
	msg.FilterType = strings.ToLower(msg.FilterType)

	return
}


