package services

import "database/sql"

type PostgresTransactionDbHandler struct {
	Db *sql.DB
}
