package repository

import "ev-events-ms/models"

func GetEventsByStatus(status string) (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Where("status = ?",status).Find(&events).Error

	return
}
func GetEventsByOwnerType(owner string) (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Where("owner_type = ?",owner).Find(&events).Error

	return
}
func GetUpcomingEvents( date string) (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Where("event_start_date >= ?", date).Find(&events).Error

	return
}
func GetEventsByDate() (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Find(&events).Error

	return
}
