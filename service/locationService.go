package service

import (
	"encoding/json"
	"ev-events-ms/models"
	"ev-events-ms/repository"
	u "ev-events-ms/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetLocations(w http.ResponseWriter, r *http.Request) () {
	location,err := repository.GetLocations()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, location)
}

func GetLocationById(w http.ResponseWriter, r *http.Request) () {
	id,err := strconv.Atoi(mux.Vars(r)["id"]) // get ID from url request
	if err != nil || id<1 { //ID is not valid
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid Id")) // REVISAR ERROR AL IMPLEMENTAR HASH
		return
	}
	location,dbErr := repository.GetLocationById(id)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Location not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, location)
}

func CreateLocation (w http.ResponseWriter, r *http.Request) {
	location := &models.Location{}
	decErr := json.NewDecoder(r.Body).Decode(location) //decode the request body into struct and failed if any error occur
	if decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	//location.Validate()

	creationErr := repository.CreateLocation(location)
	if creationErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Problem"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, location)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) ()  {
	id,err := strconv.Atoi(mux.Vars(r)["id"]) // getting id from url request
	if err != nil || id<1 { //User not found!
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid Id"))
		return
	}
	delErr1,delErr2 := repository.DeleteLocation(id)
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
	resp := u.Message(true, "Location deleted successfully")
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

func EditLocation(w http.ResponseWriter, r *http.Request) () {
	location := &models.Location{}
	editedLocation := &models.Location{}
	decodeErr := json.NewDecoder(r.Body).Decode(editedLocation) //decode the request body into struct and failed if any error occur
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	id,idErr := strconv.Atoi(mux.Vars(r)["id"]) // getting id from url request  and failed if id is not int
	if idErr != nil || id<1 {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid Id"))
		return
	}
	location, dbErr := repository.GetLocationById(id) // search user in db and failed if it doesn't exist
	if dbErr != nil {                           //User not found!
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	updateLocationValues(location,editedLocation) // update the values

	dbEditErr := repository.EditLocation(location)
	if dbEditErr != nil {
		fmt.Println(dbEditErr)
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	resp := u.Message(true, "Location deleted successfully")
	resp["location"] = location
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

var updateLocationValues = func (location *models.Location,editedLocation *models.Location)  {
	location.LocationType = editedLocation.LocationType
	location.Latitude = editedLocation.Latitude
	location.Longitude = editedLocation.Longitude
}
