package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func beginRegistration(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin registration requested")
	logRequest(r)

	user := users["username"]

	options, sessionData, err := webAuthn.BeginRegistration(user)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}
	sessionDataStore[user.Name] = sessionData
	log.Printf("sessionData=%+v\n", sessionData)

	json.NewEncoder(w).Encode(options)
	log.Println("Begin registration processed")

}

func finishRegistration(w http.ResponseWriter, r *http.Request) {
	log.Println("Finish registration requested")
	logRequest(r)

	user := users["username"]

	// Parse registration response
	session, ok := sessionDataStore[user.Name]
	if !ok {
		log.Println("Session data not found")
		return
	}

	log.Printf("session: %+v", session)
	credential, err := webAuthn.FinishRegistration(user, *session, r)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
	}
	log.Printf("err: %+v", err)
	log.Printf("credential: %+v", credential)

	user.Credentials = append(user.Credentials, *credential)
	users[user.Name] = user

	w.Write([]byte("Registration successful"))
	log.Println("Finish registration processed")
}
