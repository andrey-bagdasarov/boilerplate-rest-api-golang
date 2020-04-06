package main

import (
	"rest-api/config"
	"rest-api/services"
)

func main() {
	services.StartWebServer(config.PORT)
}
