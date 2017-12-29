package main

import (
	"log"
	"net/http"
	"service/handler"
	"os"
)

func main() {
	log.Print("Starting the service...")

	port := os.
	router := handler.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
