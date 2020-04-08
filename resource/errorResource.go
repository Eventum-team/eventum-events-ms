package resource

import "ev-events-ms/service"


func HandleErrorRequests()  {
	router := GetRouter()
	router.NotFoundHandler = service.NotFoundHandler() 	// route not found
}
