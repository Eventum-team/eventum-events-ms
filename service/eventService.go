package service

import (
	"encoding/json"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetEvents(w http.ResponseWriter, r *http.Request) () {
	events,err := repository.GetEvents()
	if err != nil {
		u.Error(w, err,http.StatusInternalServerError)
		return
	}
	u.Respond(w, events,http.StatusOK)

}

func GetEventById(w http.ResponseWriter, r *http.Request) () {
	id, idErr := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if idErr != nil {
		u.Error(w, idErr,http.StatusBadRequest)
		return
	}
	event,dbErr := repository.GetEventById(id)
	if dbErr != nil {
		u.Error(w, dbErr,http.StatusNotFound)
		return
	}
	u.Respond(w, event,http.StatusOK)

}

func CreateEvent (w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	decErr := json.NewDecoder(r.Body).Decode(event) //decode the request body
	if decErr != nil {
		u.Error(w, decErr,http.StatusBadRequest)
		return
	}
	event.AddInitialStatus() // set status to active
	valErr := event.Validate() // validate fields
	if valErr != nil{
		u.Error(w, valErr,http.StatusBadRequest)
		return
	}
	errDates := models.ProperDates(event.EventStartDate,event.EventFinishDate)
	if errDates != nil{
		u.Error(w, errDates,http.StatusBadRequest)
		return
	}
	creationErr := repository.CreateEvent(event)
	if creationErr != nil {
		u.Error(w, creationErr,http.StatusInternalServerError)
		return
	}
	u.Respond(w,u.Message( "Event Created Successfully",http.StatusCreated),http.StatusCreated)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) ()  {
	id, idErr := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if idErr != nil {
		u.Error(w, idErr,http.StatusBadRequest)
	}
	delErr1,delErr2 := repository.DeleteEvent(id)
	if delErr1 != nil {
		u.Error(w, delErr1,http.StatusNotFound)
		return
	}
	if delErr2 != nil{
		u.Error(w, delErr2,http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 doesn't allow body
	//u.Respond(w,u.Message( "Event Created Successfully"),http.StatusNoContent)
}

func EditEvent(w http.ResponseWriter, r *http.Request) () {
	event := &models.Event{}
	editedEvent := &models.Event{}
	decodeErr := json.NewDecoder(r.Body).Decode(editedEvent) //decode the request body
	if decodeErr != nil {
		u.Error(w, decodeErr,http.StatusBadRequest)
		return
	}
	id, idErr := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if idErr != nil {
		u.Error(w, idErr,http.StatusBadRequest)
	}
	event, dbErr := repository.GetEventById(id) // search user in db and failed if it doesn't exist
	if dbErr != nil {
		u.Error(w, dbErr,http.StatusNotFound)
		return
	}
	models.UpdateEventValues(event,editedEvent) // update the values
	dbEditErr := repository.EditEvent(event)
	if dbEditErr != nil {
		u.Error(w, dbEditErr,http.StatusInternalServerError)
		return
	}
	u.Respond(w,u.Message( "Event Updated Successfully",http.StatusOK),http.StatusOK)

}




