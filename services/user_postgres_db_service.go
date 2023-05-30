package services

import (
	"database/sql"
	"labora-wallet/db"
	"labora-wallet/models"
)

type PostgresUserDbHandler struct {
	Db *sql.DB
}

// Function to create an User in PostgreSQL database
func (p *PostgresUserDbHandler) CreateUser(newUser *models.User) error {
	var err error

	err = models.ValidateUser(newUser)
	if err != nil {
		return err
	}

	stmt, err := db.DbConn.Prepare("INSERT INTO public.users(user_name, document_number, document_type, country) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.UserName, newUser.DocumentNumber, newUser.DocumentType, newUser.Country)
	if err != nil {
		return err
	}

	return nil
}

// func (p *PostgresUserDbHandler) GetUser(id int) (models.User, error) {
// // Function to get the User info in PostgreSQL database

// 	return models.User{}, nil
// }

// func (p *PostgresUserDbHandler) UpdateUser(user models.User) error {
// 	// Function to update an User in PostgreSQL database
// 	return nil
// }

// func (p *PostgresUserDbHandler) DeleteUser(id int) error {
// 	// Function to delete an User in PostgreSQL database

// 	return nil
// }
