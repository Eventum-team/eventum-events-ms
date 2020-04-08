package service

import (
	"encoding/json"
	"ev-events-ms/models"
	u "ev-events-ms/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetEvents(w http.ResponseWriter, r *http.Request) () {
	events := make([] *models.Event, 0)

	err := models.GetDB().Table("events").Find(&events).Error
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	resp := u.Message(true, "success")
	resp["events"] = events
	u.Respond(w, resp)

}

func GetEventById(w http.ResponseWriter, r *http.Request) () {
	event := &models.Event{}

	id,err := strconv.Atoi(mux.Vars(r)["id"]) // get ID from url request
	if err != nil || id<1 { //ID is not valid
		u.Respond(w, u.Message(false, "Invalid Id")) // REVISAR ERROR AL IMPLEMENTAR HASH
		return
	}

	dbErr := models.GetDB().First(event,id).Error // search in
	if dbErr != nil {                                 //User not found!
		u.Respond(w, u.Message(false, "User not found"))
		return
	}
	resp := u.Message(true, "success")
	resp["event"] = event
	u.Respond(w, resp)
}

var CreateEvent = func(w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}

	err := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	//VALIDACION DE CAMPOS

	resp := event.Create() //Create account
	u.Respond(w, resp)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) ()  {
	event := &models.Event{}

	id,err := strconv.Atoi(mux.Vars(r)["id"]) // getting id from url request
	if err != nil || id<1 { //User not found!
		u.Respond(w, u.Message(false, "Invalid Id"))
		return
	}

	dbErr := models.GetDB().First(event,id).Error // search user in db and failed if it doesn't exist
	if dbErr != nil {                                 //User not found!
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	dbEditErr := models.GetDB().Delete(&event).Error
	if dbEditErr != nil {
		fmt.Println(dbEditErr)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}

func EditEvent(w http.ResponseWriter, r *http.Request) () {
	event := &models.Event{}
	editedEvent := &models.Event{}

	decodeErr := json.NewDecoder(r.Body).Decode(editedEvent) //decode the request body into struct and failed if any error occur
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	id,idErr := strconv.Atoi(mux.Vars(r)["id"]) // getting id from url request  and failed if id is not int
	if idErr != nil || id<1 {
		u.Respond(w, u.Message(false, "Invalid Id"))
		return
	}

	dbErr := models.GetDB().First(event,id).Error // search user in db and failed if it doesn't exist
	if dbErr != nil {                                 //User not found!
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	updateValues(event,editedEvent) // update the values

	dbEditErr := models.GetDB().Save(&event).Error
	if dbEditErr != nil {
		fmt.Println(dbEditErr)
		u.Respond(w, u.Message(false, "Database Connection error"))
		return
	}
	resp := u.Message(true, "success")
	resp["event"] = event
	u.Respond(w, resp)
}

func updateValues(event *models.Event,editedEvent *models.Event)  {
	event.Description = editedEvent.Description
	event.Name = editedEvent.Name
	event.EventStartDate = editedEvent.EventStartDate
	event.EventFinishDate = editedEvent.EventFinishDate
	event.Status = editedEvent.Status
	event.Url = editedEvent.Url
}

