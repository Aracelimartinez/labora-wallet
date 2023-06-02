package services

import (
	"database/sql"
	"errors"
	"labora-wallet/models"
)

type PostgresUserDbHandler struct {
	Db *sql.DB
}

var errUserNoMatch = errors.New("usuário no encontrado: Este id no existe")

// Function to create an User in PostgreSQL database
func (p *PostgresUserDbHandler) CreateUser(newUser *models.User) error {
	var err error

	err = models.ValidateUser(newUser)
	if err != nil {
		return err
	}

	stmt, err := p.Db.Prepare("INSERT INTO public.users(user_name, document_number, document_type, country, date_of_birth) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.UserName, newUser.DocumentNumber, newUser.DocumentType, newUser.Country, newUser.DateOfBirth)
	if err != nil {
		return err
	}

	return nil
}

// Function to get the User info in PostgreSQL database
func (p *PostgresUserDbHandler) GetUser(id int) (*models.User, error) {

	var err error
	var user models.User

	stmt, err := p.Db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.UserName, &user.DocumentNumber, &user.DocumentType, &user.Country, &user.DateOfBirth, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errUserNoMatch
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// func (p *PostgresUserDbHandler) UpdateUser(user models.User) error {
// 	// Function to update an User in PostgreSQL database
// 	return nil
// }

// func (p *PostgresUserDbHandler) DeleteUser(id int) error {
// 	// Function to delete an User in PostgreSQL database

// 	return nil
// }
