package repository

import (
	"ev-events-ms/models"
	"time"
)

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
func GetEventsByRangeDate(date1 time.Time, date2 time.Time) (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	err = GetDB().Table("events").Where("event_start_date >= ? AND event_start_date <= ? ", date1, date2).Find(&events).Error
	return
}
func GetEventsByName(name string) (events [] *models.Event,err error)  {
	events = make([]*models.Event, 0)
	str := "%" + name + "%"
	err = GetDB().Table("events").Where("name LIKE ?", str).Find(&events).Error
	return
}
