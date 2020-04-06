package services

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting server on port " + port)
	var router = InitRouter()
	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("Server starting error")
		log.Println("Error..." + err.Error())
	}
}
