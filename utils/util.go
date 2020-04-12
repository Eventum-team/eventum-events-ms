package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func Message(message string, status int) map[string]interface{} {
	return map[string]interface{}{ "message": message, "status": status}
}
func ErrorMessage(err error, status int) map[string]interface{} {
	return map[string]interface{}{ "message": err.Error(), "status": status}
}

func Respond(w http.ResponseWriter, data interface{},statusCode int ) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil{
		fmt.Println(err)
		//Error(w,err)
		return
	}
}

func Error(w http.ResponseWriter, err error,statusCode int) {
	if err != nil {
		Respond(w, ErrorMessage(err,statusCode),statusCode)
		return
	}

}

