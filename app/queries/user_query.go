package queries

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserById(id uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := q.Get(&user, query, id)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE email = $1`
	err := q.Get(&user, query, email)

	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Send query to database.
	_, err := q.Exec(
		query, 
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.Firstname, u.Lastname, "_blank", u.PasswordHash, u.UserStatus, u.UserRole,
	)

	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
