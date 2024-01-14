package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

func beginRegistration(w http.ResponseWriter, r *http.Request) {
	user := &User{
		ID:          []byte("user-id"),
		Name:        "username",
		DisplayName: "User Name",
		Icon:        "",
	}

	options, sessionData, err := webAuthn.BeginRegistration(user)
	if err != nil {
		fmt.Println(err)
	}
	sessionDataStore[user.Name] = sessionData

	json.NewEncoder(w).Encode(options)
}

func finishRegistration(w http.ResponseWriter, r *http.Request) {
	user := &User{
		ID:          []byte("user-id"),
		Name:        "username",
		DisplayName: "User Name",
		Icon:        "",
		Credentials: []webauthn.Credential{},
	}

	// Parse registration response
	session, ok := sessionDataStore[user.Name]
	if !ok {
		fmt.Println("Session data not found")
		return
	}

	credential, err := webAuthn.FinishRegistration(user, *session, r)
	if err != nil {
		fmt.Println(err)
	}

	user.Credentials = append(user.Credentials, *credential)
	users[user.Name] = user

	w.Write([]byte("Registration successful"))
}
