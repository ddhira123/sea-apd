package postgres

import (
	"database/sql"
)

type postgresProductRepository struct {
	db *sql.DB
}
