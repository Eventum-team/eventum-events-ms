
package main

import (
	"ev-events-ms/resource"
	"fmt"
	"net/http"
	"os"
)

func main() {

	resource.HandleEventRequests()
	resource.HandleEventFilterRequests()
	resource.HandleErrorRequests()

	port := os.Getenv("PORT") //setting port according to .env info
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, resource.GetRouter()) //Launching app
	if err != nil {
		fmt.Print(err)
	}
}
