package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/model"
)

func PersonHandler(w http.ResponseWriter, req *http.Request) {
	address := model.Address{"Sunset blv", "San Diego", "California"}
	person := model.Person{"Hill", "Williams", address}

	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Println("JSON serialization error " + err.Error())
	}
}
