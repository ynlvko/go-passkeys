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
	RPID:          "go-passkeys.onrender.com",
	RPDisplayName: "Demo App",
	RPOrigins:     []string{"https://go-passkeys.onrender.com"},
})

func main() {
	http.HandleFunc("/begin-registration", beginRegistration)
	http.HandleFunc("/finish-registration", finishRegistration)

	http.ListenAndServe("0.0.0.0", nil)
}
