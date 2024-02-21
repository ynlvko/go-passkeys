package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func beginLogin(w http.ResponseWriter, r *http.Request) {
	user, ok := users["username"]
	if !ok {
		log.Printf("User not found")
	}

	options, sessionData, err := webAuthn.BeginLogin(user)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}
	sessionDataStore[user.Name] = sessionData
	log.Printf("sessionData=%+v\n", sessionData)

	json.NewEncoder(w).Encode(options)
	log.Println("Begin registration processed")
}
