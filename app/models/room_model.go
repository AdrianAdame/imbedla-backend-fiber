package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	UserId    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Name      string    `db:"name" json:"name" validate:"required,lte=255"`
	Color     string    `db:"color" json:"color" validate:"required,lte=255"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Type      string    `db:"type" json:"type" validate:"required,lte=255"`
}