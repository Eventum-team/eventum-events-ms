package models

import (
	u "ev-events-ms/utils"
	"github.com/jinzhu/gorm"

	//u "../utils"
)
type Location struct {
	gorm.Model
	locationId string
	locationType string
	latitude string
	longitude string
	eventId string
}

func (location *Location) validate() (map[string]interface{}, bool) {

	return u.Message(false, "Requirement passed"), true
}

func (location *Location) Create() (map[string] interface{}) {

	if resp, ok := location.validate(); !ok {
		return resp
	}

	GetDB().Create(location)

	//if event.ID <= 0 {
	//	return u.Message(false, "Failed to create account, connection error.")
	//}

	response := u.Message(true, "Location added succesfully")
	response["location"] = location
	return response
}