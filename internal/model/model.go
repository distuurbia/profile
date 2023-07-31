package model

import "github.com/google/uuid"

type Profile struct {
	Age int
	Password []byte
	RefreshToken []byte
	Username string
	Country string
	ID uuid.UUID
}