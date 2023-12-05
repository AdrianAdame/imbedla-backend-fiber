package queries

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/jmoiron/sqlx"
)

type TokenQueries struct {
	*sqlx.DB
}

func (q *TokenQueries) CreateTokens(t *models.Tokens) error {
	query := `INSERT INTO tokens VALUES ($1,$2,$3,$4,$5)`

	_, err := q.Exec(
		query,
		t.TokenID, t.UserID, t.Access, t.Refresh, t.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (q *TokenQueries) DeleteTokens(t *models.Tokens) error {

	query := `DELETE FROM tokens WHERE id = $1;`

	_, err := q.Exec(
		query,
		t.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}
