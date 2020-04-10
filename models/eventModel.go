package models

import (
	"errors"
	u "ev-events-ms/utils"
	"time"
)

type Event struct {
	ID string `json:"id"`
	Counter uint64 `gorm:"auto_increment"`
	OwnerId string `gorm:"not null" json:"ownerId"`
	OwnerType string ` gorm:"not null"json:"ownerType"`
	Name string `gorm:"not null" json:"name"`
	Status string `gorm:"not null" json:"status"`
	EventStartDate time.Time `json:"eventStartDate"`
	EventFinishDate time.Time `json:"eventFinishDate"`
	Description string `json:"description"`
	Url string `json:"url"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`

}

var StatusOptions = []string{
	"active",
	"cancelled",
	"ended",
}
var OwnerTypeOptions = []string{
	"group",
	"user",
}

func (event *Event) Validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}

func (event *Event) AddInitialStatus() {
	event.Status = StatusOptions[0]
}

func (event *Event) ProperDates() error  {
	if event.EventStartDate.After(event.EventFinishDate){
		return errors.New("Start date is later than finish date")
	}
	return nil
}

var UpdateEventValues = func (event *Event,editedEvent *Event)  {
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
func ValidateDate(date time.Time) (valid bool){
	valid = true

	return
}



