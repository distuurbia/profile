// Package model contains models of project
package model

import "github.com/google/uuid"

//Profile contains fields that we have in our postgresql table profiles
type Profile struct {
	Age int
	Password []byte
	RefreshToken []byte
	Username string
	Country string
	ID uuid.UUID
}

// TokenPair contains an access and a refresh tokens
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}