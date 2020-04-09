package service

import (
	"encoding/json"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) () {
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

func GetEventById(w http.ResponseWriter, r *http.Request) () {
	id := mux.Vars(r)["id"] // get ID from url request
	event,dbErr := repository.GetEventById(id)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "User not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, event)
}

func CreateEvent (w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	decErr := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	if decErr != nil {
		fmt.Println(decErr)
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	_, createdErr := repository.GetEventById(event.ID) // check if event already exists
	if createdErr == nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Event already Exists"))
		return
	}
	//event.Validate()
	creationErr := repository.CreateEvent(event)
	if creationErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, creationErr.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) ()  {
	id := mux.Vars(r)["id"] // getting id from url request

	delErr1,delErr2 := repository.DeleteEvent(id)
	if delErr1 != nil {
		w.WriteHeader(http.StatusNotFound)//User not found!
		u.Respond(w, u.Message(false, "User not found"))
		return
	}
	if delErr2 != nil{
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	resp := u.Message(true, "Event deleted successfully")
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

func EditEvent(w http.ResponseWriter, r *http.Request) () {
	event := &models.Event{}
	editedEvent := &models.Event{}
	decodeErr := json.NewDecoder(r.Body).Decode(editedEvent) //decode the request body into struct and failed if any error occur
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	id := mux.Vars(r)["id"]
	event, dbErr := repository.GetEventById(id) // search user in db and failed if it doesn't exist
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	updateEventValues(event,editedEvent) // update the values

	dbEditErr := repository.EditEvent(event)
	if dbEditErr != nil {
		fmt.Println(dbEditErr)
		w.WriteHeader(http.StatusInternalServerError)

		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	u.Respond(w, event)
	w.Header().Set("Content-Type", "application/json")
}

var updateEventValues = func (event *models.Event,editedEvent *models.Event)  {
	//Possible changes
	event.Description = editedEvent.Description
	event.Name = editedEvent.Name
	event.EventStartDate = editedEvent.EventStartDate
	event.EventFinishDate = editedEvent.EventFinishDate
	event.Status = editedEvent.Status
	event.Url = editedEvent.Url
	event.Longitude = editedEvent.Longitude
	event.Latitude= editedEvent.Latitude
}


