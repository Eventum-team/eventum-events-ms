package models

import (
	"github.com/jinzhu/gorm"
	u "ev-events-ms/utils"
)

type Event struct {
	gorm.Model
	EventId string
	OwnerId string
	OwnerType string
	Description string
	Name string
	EventStartDate string
	EventFinishDate string
	Status string
	Url string
	Location Location `gorm:"foreignkey:EventId"`
}

func (event *Event) validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}

func (event *Event) Create() (map[string] interface{}) {

	GetDB().Create(event)

	//if event.ID <= 0 {
	//	return u.Message(false, "Failed to create account, connection error.")
	//}

	response := u.Message(true, "Event has been created")
	response["event"] = event
	return response
}