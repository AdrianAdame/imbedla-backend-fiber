package models

import "github.com/google/uuid"

type RoomUser struct {
	ID uuid.UUID `json:"email" validate:"required, uuid"`
}