package repository

import (
	"ev-events-ms/models"
)

func GetEvents() (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Find(&events).Error

	return
}

func GetEventById(id string ) (event *models.Event,err error) {
	event = &models.Event{}

	//location, err := GetLocationByEventId(id)
	//if err != nil {
	//	fmt.Println("No location found")
	//}
	err = GetDB().First(event, id).Error

	//event.Location = models.Location{
	//	Latitude: location.Latitude,
	//	Longitude: location.Longitude,
	//	LocationType: location.LocationType,
	//	ID: location.ID,
	//	EventID: location.EventID,
	//
	//}
	return
}

func CreateEvent(event *models.Event) (err error) {
	err= GetDB().Create(event).Error
	return
}

func DeleteEvent(id int) (err1 error,err2 error)  {
	event := &models.Event{}
	err1 = GetDB().First(event,id).Error
	if err1 != nil {
		return
	}
	err2 = GetDB().Delete(&event).Error
	return
}

func EditEvent(editedEvent *models.Event) (err error)  {
	// fields updated in eventService
	err = GetDB().Save(&editedEvent).Error
	return
}