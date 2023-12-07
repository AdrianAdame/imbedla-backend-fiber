package models

import "github.com/google/uuid"

type RoomUser struct {
	ID uuid.UUID `json:"user_id" validate:"required, uuid"`
}