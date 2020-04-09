package utils

import (
	"encoding/json"
	"net/http"
)


func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
}



func Respond(w http.ResponseWriter, data interface{} ) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil{
		print("Problems encoding data to Json format")
	}
}

