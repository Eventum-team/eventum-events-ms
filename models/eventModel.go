package models

import (
	"errors"
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

func (event *Event) Validate() (err error) {
	e := ""
	if event.Name == ""{
	//return  errors.New("Invalid fields")
		e+= "name"
	}
	if event.OwnerType == ""{
		e+= ", ownerType"
	}
	if event.OwnerId == ""{
		e+= ", OwnerId"
	}
	if !ValidateDate(&event.EventStartDate){
		e+= ", eventStartDate"
	}
	if !ValidateDate(&event.EventFinishDate){
		e+= ", eventFinishDate"
	}
	if e != ""{
		return errors.New("invalid Fields: "+ e)
	}
	return nil
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
func ValidateDate(date *time.Time) (valid bool){
	errDate := "0001-01-01 00:00:00 +0000 UTC"
	return date.String() != errDate
}



