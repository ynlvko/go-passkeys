package main

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"log"
	"net/http"
	"os"
)

// In-memory user storage
var users = map[string]*User{}
var sessionDataStore = map[string]*webauthn.SessionData{}

// WebAuthn configuration
var webAuthn, _ = webauthn.New(&webauthn.Config{
	RPID:          "go-passkeys.onrender.com",
	RPDisplayName: "Demo App",
	RPOrigins:     []string{"https://go-passkeys.onrender.com"},
})

func main() {
	http.HandleFunc("/begin-registration", beginRegistration)
	http.HandleFunc("/finish-registration", finishRegistration)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
