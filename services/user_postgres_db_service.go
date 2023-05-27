package services

import (
	"database/sql"
	"labora-wallet/models"
)

type PostgresUserDbHandler struct {
	Db *sql.DB
}

func (p *PostgresUserDbHandler) CreateUser(user models.User) error {
	// Implementar la lógica para crear un user en la base de datos PostgreSQL

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
