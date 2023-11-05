package model

import "github.com/google/uuid"

type AuthLogin struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
