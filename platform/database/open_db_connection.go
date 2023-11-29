package database

import (
	"fmt"
	"os"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/queries"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
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

	fmt.Println(dbType)

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
		UserQueries: &queries.UserQueries{DB: db},
	}, nil
}
