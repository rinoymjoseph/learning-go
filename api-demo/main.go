package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("Creating routes")
	//specify endpoints
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/warriors", getWarriors).Methods("GET")
	http.Handle("/", router)
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
