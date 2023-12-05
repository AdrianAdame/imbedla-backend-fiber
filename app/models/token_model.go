package models

import (
	"time"

	"github.com/google/uuid"
)

// Renew struct to describe refresh token object.
type Renew struct {
	RefreshToken string `json:"refresh_token"`
}

type Tokens struct {
	TokenID   uuid.UUID `db:"token_id" json:"token_id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Access    string    `db:"access_token" json:"access_token"`
	Refresh   string    `db:"refresh_token" json:"refresh_token"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
