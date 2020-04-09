package models

import (
	u "ev-events-ms/utils"
)

type Event struct {
	ID string
	Counter uint64 `gorm:"auto_increment"`
	OwnerId string
	OwnerType string
	Description string
	Name string
	EventStartDate string
	EventFinishDate string
	Status string
	Url string
	Latitude string
	Longitude string

}


func (event *Event) Validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}

