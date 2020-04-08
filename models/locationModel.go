package models

import (
	u "ev-events-ms/utils"
)
type Location struct {
	//gorm.Model
	ID uint64 `gorm:"type: serial;"`
	LocationType string
	Latitude string
	Longitude string
	EventId string
}

func (location *Location) validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}
