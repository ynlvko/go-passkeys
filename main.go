package main

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// In-memory user storage
var users = map[string]*User{}
var sessionDataStore = map[string]*webauthn.SessionData{}

// WebAuthn configuration
var webAuthn *webauthn.WebAuthn

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	var err error
	webAuthn, err = webauthn.New(&webauthn.Config{
		RPID:          os.Getenv("RPID"),
		RPDisplayName: os.Getenv("RPDisplayName"),
		RPOrigins:     []string{os.Getenv("RPOrigin"), "android:apk-key-hash:r2suYH6W-MNw2jid65geNrmF-ef0BcNzZjaAw389Dn8"},
	})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/begin-registration", corsMiddleware(http.HandlerFunc(beginRegistration)))
	http.Handle("/finish-registration", corsMiddleware(http.HandlerFunc(finishRegistration)))

	http.HandleFunc("/.well-known/apple-app-site-association", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "static/apple-app-site-association")
	})
	http.HandleFunc("/.well-known/assetlinks.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "static/assetlinks")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
