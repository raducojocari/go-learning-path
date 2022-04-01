package main

import (
	"api2/controllers"
	"log"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3001", nil)

	if err != nil {
		log.Fatal(err)
	}

}
