package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
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

func Error(w http.ResponseWriter, err error) {
	if err != nil {
		Respond(w, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),

		})
		return
	}

}

