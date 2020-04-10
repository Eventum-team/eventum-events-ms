package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func Message(message string) map[string]interface{} {
	return map[string]interface{}{ "message": message}
}

func Respond(w http.ResponseWriter, data interface{} ) {
	err := json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
	if err != nil{
		fmt.Println(err)
		//Error(w,err)
		return
	}
}

func Error(w http.ResponseWriter, err error,statusCode int,des string) {
	if err != nil {
		Respond(w, struct {
			Error string `json:"error"`
			Code int `json:"error"`
			Description string `json:"error"`
		}{
			Error: err.Error(),
			Code: statusCode,
			Description: des,

		})
		return
	}

}

