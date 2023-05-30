package services

import (
	"database/sql"
	"labora-wallet/db"
	"labora-wallet/models"
)

type PostgresLogDbHandler struct {
	Db *sql.DB
}

// Function to create a Log in PostgreSQL database
func (p *PostgresLogDbHandler) CreateLog(User *models.User, canCreate bool) (models.Log, error) {
	var err error

	stmt, err := db.DbConn.Prepare("INSERT INTO public.logs(approved, user_id) VALUES ($1, $2) RETURNING id, approved, user_id, created_at")
	if err != nil {
		return models.Log{}, err
	}

	defer stmt.Close()

	var createdLog models.Log
	err = stmt.QueryRow(canCreate, User.ID).Scan(&createdLog.ID, &createdLog.Approved, &createdLog.UserID, &createdLog.CreatedAt)
	if err != nil {
		return models.Log{}, err
	}

	return createdLog, nil
}
