package models

import (
	u "ev-events-ms/utils"
)

type Event struct {
	//gorm.Model
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
	Location []Location `gorm:"foreignkey:EventID"`

}

func (event *Event) Validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}

