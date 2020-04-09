package repository

import (
	"ev-events-ms/models"
)

func GetLocations() (locations [] *models.Event,err error)  {
	locations = make([]*models.Event, 0)
	err = GetDB().Table("events").Select("latitude, longitude, id, name").Find(&locations).Error
	return
}


