package database

import (
	"os"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/queries"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
	*queries.TokenQueries
	*queries.RoomQueries
	*queries.PlantQueries
}

// type DBInstances struct {
// 	Database []Queries
// }

func OpenDBConnection() (*Queries, error) {
	var (
		db  *sqlx.DB
		err error
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "pgx":
		db, err = PostgresSQLConnection()
	}

	if err != nil {
		return nil, err
	}

	// databases := []Queries{
	// 	{UserQueries: &queries.UserQueries{DB: db}},
	// }

	// return &DBInstances{
	// 	Database: databases,
	// }, nil

	return &Queries{
		UserQueries:  &queries.UserQueries{DB: db},
		TokenQueries: &queries.TokenQueries{DB: db},
		RoomQueries: &queries.RoomQueries{DB: db},
		PlantQueries: &queries.PlantQueries{DB: db},
	}, nil
}
