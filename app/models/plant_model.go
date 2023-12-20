package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PlantH struct {
	ID                uuid.UUID       `db:"id" json:"id"`
	UserId            uuid.UUID       `db:"user_id" json:"user_id" validate:"required,uuid"`
	RoomId            uuid.UUID       `db:"room_id" json:"room_id" validate:"required,uuid"`
	Name              string          `db:"name" json:"name" validate:"required,lte=255"`
	RefPlant          string          `db:"ref_plant" json:"ref_plant" validate:"required,lte=255"`
	CreatedAt         time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at" json:"updated_at"`
	ModuleInformation json.RawMessage `db:"module_information" json:"module_information"`
	ModuleSpecs       json.RawMessage `db:"module_specs" json:"module_specs"`
	Favorite          bool            `db:"favorite" json:"favorite"`
}

type PlantD struct {
	ID                uuid.UUID `db:"id" json:"id"`
	UserId            uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	RoomId            uuid.UUID `db:"room_id" json:"room_id" validate:"required,uuid"`
	Name              string    `db:"name" json:"name" validate:"required,lte=255"`
	RefPlant          string    `db:"ref_plant" json:"ref_plant" validate:"required,lte=255"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
	ModuleInformation string    `db:"module_information" json:"module_information"`
	ModuleSpecs       string    `db:"module_specs" json:"module_specs"`
	Favorite          bool      `db:"favorite" json:"favorite"`
}

type DeletePlant struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type UpdatePlantH struct {
	ID                uuid.UUID       `json:"id" validate:"required,uuid"`
	Name              string          `json:"name" validate:"lte=255"`
	ModuleInformation json.RawMessage `json:"module_information"`
	ModuleSpecs       json.RawMessage `json:"module_specs"`
	Favorite          bool            `json:"favorite"`
}
