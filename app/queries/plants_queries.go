package queries

import (
	"fmt"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PlantQueries struct {
	*sqlx.DB
}

func (q *PlantQueries) GetAllPlantsByRoomId(roomId uuid.UUID) ([]models.PlantD, error) {
	plants := []models.PlantD{}
	query := `SELECT * FROM plants WHERE room_id = $1 ORDER BY updated_at DESC`
	err := q.Select(&plants, query, roomId)

	if err != nil {
		return plants, err
	}

	return plants, nil
}

func (q *PlantQueries) GetPlantById(id uuid.UUID) (models.PlantD, error) {
	plant := models.PlantD{}
	query := `SELECT * FROM plants WHERE id = $1`
	err := q.Get(&plant, query, id)

	if err != nil {
		return plant, err
	}

	return plant, nil
}

func (q *PlantQueries) CreatePlant(p *models.PlantD) error {
	query := `INSERT INTO plants VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := q.Exec(
		query,
		p.ID, p.UserId, p.RoomId, p.Name, p.RefPlant, p.CreatedAt, p.UpdatedAt, p.ModuleInformation, p.ModuleSpecs, p.Favorite,
	)

	fmt.Println(p)

	if err != nil {
		return err
	}

	return nil
}

func (q *PlantQueries) EditPlant(r *models.PlantD) error {
	query := `UPDATE plants SET name = $1, module_information = $2, module_specs = $3, updated_at = $4 favorite = $5 WHERE id = $6`

	_, err := q.Exec(
		query,
		r.Name,
		r.ModuleInformation,
		r.ModuleSpecs,
		r.UpdatedAt,
		r.Favorite,
		r.ID,
	)

	if err != nil {
		return err
	}

	return err
}

func (q *RoomQueries) DeletePlant(id uuid.UUID) error {
	query := `DELETE FROM plants WHERE id = $1`

	_, err := q.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
