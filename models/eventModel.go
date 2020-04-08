package models

import (
	u "ev-events-ms/utils"
	"github.com/jinzhu/gorm"
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

func (event *Event) Validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}

