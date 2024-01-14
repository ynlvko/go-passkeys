package main

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

// In-memory user storage
var users = map[string]*User{}
var sessionDataStore = map[string]*webauthn.SessionData{}

// WebAuthn configuration
var webAuthn, _ = webauthn.New(&webauthn.Config{
	RPID:          "localhost",
	RPDisplayName: "Demo App",
	RPOrigins:     []string{"http://localhost:8080"},
})

func main() {
	http.HandleFunc("/begin-registration", beginRegistration)
	http.HandleFunc("/finish-registration", finishRegistration)

	http.ListenAndServe(":8080", nil)
}
