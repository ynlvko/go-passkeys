package main

import "github.com/go-webauthn/webauthn/webauthn"

// User represents a basic user model
type User struct {
	ID          []byte
	Name        string
	DisplayName string
	Icon        string
	Credentials []webauthn.Credential
}

// Implementing webauthn.User interface
func (u *User) WebAuthnID() []byte {
	return u.ID
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnIcon() string {
	return u.Icon
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}
