package queries

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoomQueries struct {
	*sqlx.DB
}

func (q *RoomQueries) GetRoomById(id uuid.UUID) (models.Room, error) {
	room := models.Room{}
	query := `SELECT * FROM rooms WHERE id = $1`
	err := q.Get(&room, query, id)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (q *RoomQueries) GetRoomsByUserId(userId uuid.UUID) ([]models.Room, error) {
	rooms := []models.Room{}
	query := `SELECT * FROM rooms WHERE user_id = $1`
	err := q.Select(&rooms, query, userId)

	if err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (q *RoomQueries) CreateRoom(r *models.Room) error {
	query := `INSERT INTO rooms VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := q.Exec(
		query,
		r.ID, r.UserId, r.Name, r.Color, r.CreatedAt, r.UpdatedAt, r.Type,
	)

	if err != nil {
		return err
	}


	return nil
}

