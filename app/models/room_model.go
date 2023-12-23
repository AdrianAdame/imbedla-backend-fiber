package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        		uuid.UUID `db:"id" json:"id"`
	UserId    		uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Name      		string    `db:"name" json:"name" validate:"required,lte=255"`
	Color     		string    `db:"color" json:"color" validate:"required,lte=255"`
	CreatedAt 		time.Time `db:"created_at" json:"created_at"`
	UpdatedAt 		time.Time `db:"updated_at" json:"updated_at"`
	Type      		string    `db:"type" json:"type" validate:"required,lte=255"`
	TotalPlants     string    `db:"total_plants" json:"total_plants"`
	LatestPlantName string    `db:"latest_plant_name" json:"latest_plant_name"`
}

type DeleteRoom struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type UpdateRoom struct {
	ID    uuid.UUID `json:"id" validate:"required,uuid"`
	Name  string    `json:"name" validate:"lte=255"`
	Color string    `json:"color" validate:"lte=255"`
	Type  string	`json:"type" validate:"lte=255"`
}