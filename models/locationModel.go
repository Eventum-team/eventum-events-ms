package models

import (
	u "ev-events-ms/utils"
)
type Location struct {
	ID uint64 `gorm:"auto_increment;primary_key"`
	Latitude string
	Longitude string
	EventID string
}

func (location *Location) validate() (map[string]interface{}, bool) {
	return u.Message(false, "Requirement passed"), true
}
