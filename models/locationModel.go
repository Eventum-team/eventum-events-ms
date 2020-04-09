package models

import (
	u "ev-events-ms/utils"
)
type Location struct {
	//gorm.Model
	ID uint64 `gorm:"auto_increment;primary_key"`
	LocationType string
	Latitude string
	Longitude string
	EventID string
}

func (location *Location) validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}
