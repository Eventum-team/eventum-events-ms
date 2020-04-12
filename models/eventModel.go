package models

import (
	"errors"
	"time"
)

type Event struct {
	ID uint64 `gorm:"auto_increment" json:"id"`
	OwnerId uint64 `gorm:"not null" json:"ownerId"`
	OwnerType string ` gorm:"not null"json:"ownerType"`
	Name string `gorm:"not null" json:"name"`
	Status string `gorm:"not null" json:"status"`
	EventType string `gorm:"not null" json:"eventType"`
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

var EventTypeOptions = []string{
	"official",
	"unofficial",
}


func (event *Event) Validate() (err error) {
	e := ""
	if event.Name == ""{
		e+= " name"
	}
	ot:= false
	for _, val :=range OwnerTypeOptions{
		if event.OwnerType == val{
			ot=true
		}
	}
	if ot == false{
		e+= "ownerType"
	}
	et:= false
	for _, val :=range EventTypeOptions{
		if event.EventType == val{
			et=true
		}
	}
	if et == false{
		e+= " eventType"
	}
	if event.OwnerId == 0 {
		e+= " OwnerId"
	}
	if !ValidateDate(&event.EventStartDate){
		e+= " eventStartDate"
	}
	if !ValidateDate(&event.EventFinishDate){
		e+= " eventFinishDate"
	}
	if e != ""{
		return errors.New("invalid Fields: "+ e)
	}
	return nil
}

func (event *Event) AddInitialStatus() {
	event.Status = StatusOptions[0]
}

func  ProperDates(d1 time.Time,d2 time.Time) error  {
	if d1.After(d2){
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
	event.EventType = editedEvent.EventType
}
func ValidateDate(date *time.Time) (valid bool){
	errDate := "0001-01-01 00:00:00 +0000 UTC"
	return date.String() != errDate
}
func ValidateStringDate(date string) (d time.Time,valid bool){
	d, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return d,false
	}
	return d,true
}



