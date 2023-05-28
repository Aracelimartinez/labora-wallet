package services

import (
	"database/sql"
	"labora-wallet/db"
	"labora-wallet/models"
)

type Use struct {
	ID             int    `json:"id"`
	UserName       string `json:"user_name"`
	DocumentNumber string `json:"document_number"`
	DocumentType   string `json:"document_type"`
	Country        string `json:"country"`
}

type PostgresUserDbHandler struct {
	Db *sql.DB
}

func (p *PostgresUserDbHandler) CreateUser(newUser models.User) error {
	// Implementar la lógica para crear un user en la base de datos PostgreSQL
	var err error

	err = models.ValidateUser(&newUser)
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
// 	// Implementar la lógica para obtener el status de un user de la base de datos PostgreSQL

// 	return models.User{}, nil
// }

// func (p *PostgresUserDbHandler) UpdateUser(user models.User) error {
// 	// Implementar la lógica para actualizar un user en la base de datos PostgreSQL
// 	return nil
// }

// func (p *PostgresUserDbHandler) DeleteUser(id int) error {
// 	// Implementar la lógica para eliminar un user de la base de datos PostgreSQL

// 	return nil
// }
