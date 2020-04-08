package repository

import "ev-events-ms/models"

func GetLocations() (locations [] *models.Location,err error)  {
	locations = make([]*models.Location, 0)
	err = GetDB().Table("locations").Find(&locations).Error
	return
}

func GetLocationById(id int) (location *models.Location,err error) {
	location = &models.Location{}
	err = GetDB().First(location, id).Error
	return
}

func CreateLocation(location *models.Location) (err error) {
	err= GetDB().Create(location).Error
	return
}

func DeleteLocation(id int) (err1 error,err2 error)  {
	location := &models.Location{}
	err1 = GetDB().First(location,id).Error
	if err1 != nil {
		return
	}
	err2 = GetDB().Delete(&location).Error
	return
}

func EditLocation(editedLocation *models.Location) (err error)  {
	// fields updated in locationService
	err = GetDB().Save(&editedLocation).Error
	return
}
