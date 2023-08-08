package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering health check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func getWarriors(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering getWarriors")
	persons := buildWarriors()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(persons)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func buildWarriors() []Warrior {
	var warriors []Warrior

	var warrior Warrior
	warrior.Id = 1
	warrior.Name = "Link"
	warriors = append(warriors, warrior)

	warrior.Id = 2
	warrior.Name = "Zelda"
	warriors = append(warriors, warrior)

	warrior.Id = 3
	warrior.Name = "Revali"
	warriors = append(warriors, warrior)
	return warriors
}
