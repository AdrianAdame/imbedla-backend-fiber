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
	query := `SELECT r.*, COUNT(p.id) as total_plants, COALESCE((SELECT name FROM plants WHERE room_id = r.id ORDER BY created_at DESC LIMIT 1), '') AS latest_plant_name FROM rooms r LEFT JOIN plants p ON r.id = p.room_id WHERE r.id = $1 GROUP BY r.id ORDER BY r.updated_at DESC`
	err := q.Get(&room, query, id)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (q *RoomQueries) GetRoomsByUserId(userId uuid.UUID) ([]models.Room, error) {
	rooms := []models.Room{}
	query := `SELECT r.*, COUNT(p.id) as total_plants, COALESCE((SELECT name FROM plants WHERE room_id = r.id ORDER BY created_at DESC LIMIT 1), '') AS latest_plant_name FROM rooms r LEFT JOIN plants p ON r.id = p.room_id WHERE r.user_id = $1 GROUP BY r.id ORDER BY r.updated_at DESC`
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

func (q *RoomQueries) EditRoom(r *models.Room) error {
	query := `UPDATE rooms SET name = $1, color = $2, updated_at = $3, type = $4 WHERE id = $5`

	_, err := q.Exec(
		query,
		r.Name,
		r.Color,
		r.UpdatedAt,
		r.Type,
		r.ID,
	)

	if err != nil {
		return err
	}

	return err
}

func (q *RoomQueries) DeleteRoom(id uuid.UUID) error {
	query := `DELETE FROM rooms WHERE id = $1`

	_, err := q.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
