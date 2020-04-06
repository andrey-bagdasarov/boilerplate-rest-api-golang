package services

import (
	"net/http"
	"rest-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{"GetPersons",
		"GET",
		"/persons",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			handlers.PersonHandler(writer, request)
		}},
}
