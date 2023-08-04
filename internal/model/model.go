// Package model contains models of project
package model

import "github.com/google/uuid"

// Profile contains fields that we have in our postgresql table profiles
type Profile struct {
	Age          int32 `validate:"gte=18,lte=120"`
	ID           uuid.UUID
	Username     string `validate:"required,min=4,max=20"`
	Country      string `validate:"required,min=2"`
	Password     []byte `validate:"required,min=4"`
	RefreshToken []byte
}
