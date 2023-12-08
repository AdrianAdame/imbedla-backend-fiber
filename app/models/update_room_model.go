package models

import "github.com/google/uuid"

type UpdateRoom struct {
	ID    uuid.UUID `json:"id" validate:"required,uuid"`
	Name  string    `json:"name" validate:"lte=255"`
	Color string    `json:"color" validate:"lte=255"`
	Type  string	`json:"type" validate:"lte=255"`
}