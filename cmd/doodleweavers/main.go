package main

import (
	"log"
	"net/http"
)

func main() {

	serverAddress := ":8080"
	log.Printf("Starting Doodle Weavers server on %s\n", serverAddress)

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal("Error Starting Server: ", err)
	}

}
