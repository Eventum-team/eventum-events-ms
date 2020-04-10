package service

import (
	"encoding/json"
	"errors"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) () {
	events,err := repository.GetEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, err)
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
		u.Error(w, dbErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, event)

}

func CreateEvent (w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	decErr := json.NewDecoder(r.Body).Decode(event) //decode the request body
	if decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Error(w, decErr)
		return
	}
	_, createdErr := repository.GetEventById(event.ID) // check if event already exists //NO DEBERIA POR QUE PASAR YO GENERO ID
	if createdErr == nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, errors.New("user already exists"))
		return
	}
	 event.AddInitialStatus() // set status to active
	//event.Validate()
	errDates := event.ProperDates()
	if errDates != nil{
		w.WriteHeader(http.StatusBadRequest)
		u.Error(w, errDates)
		return
	}
	creationErr := repository.CreateEvent(event)
	if creationErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, creationErr)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) ()  {
	id := mux.Vars(r)["id"] // getting id from url request
	delErr1,delErr2 := repository.DeleteEvent(id)
	if delErr1 != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Error(w, delErr1)
		return
	}
	if delErr2 != nil{
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, delErr2)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 doesn't allow body
}

func EditEvent(w http.ResponseWriter, r *http.Request) () {
	event := &models.Event{}
	editedEvent := &models.Event{}
	decodeErr := json.NewDecoder(r.Body).Decode(editedEvent) //decode the request body
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Error(w, decodeErr)
		return
	}
	id := mux.Vars(r)["id"]
	event, dbErr := repository.GetEventById(id) // search user in db and failed if it doesn't exist
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Error(w, dbErr)
		return
	}
	models.UpdateEventValues(event,editedEvent) // update the values
	dbEditErr := repository.EditEvent(event)
	if dbEditErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Error(w, dbEditErr)
		return
	}
	w.WriteHeader(http.StatusOK)

}




