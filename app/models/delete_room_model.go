package models

import "github.com/google/uuid"

type DeleteRoom struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}