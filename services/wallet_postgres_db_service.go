package services

import (
	"database/sql"
	"errors"
	"fmt"
	"labora-wallet/db"
	"labora-wallet/models"
)

type PostgresWalletDbHandler struct {
	Db *sql.DB
}

// type Wallet struct {
// 	ID        int     `json:"id"`
// 	Balance   float64 `json:"balance"`
// 	Currency  string  `json:"currency"`
// 	LogID     int     `json:"log_id"`
// 	CreatedAt string  `json:"created_at"`
// }

var ErrNoMatch = errors.New("Billetera no encontrada: Este id no existe")

func (p *PostgresWalletDbHandler) CreateWallet(wallet *models.Wallet) error {
	// Implementar la lógica para crear una wallet en la base de datos PostgreSQL

	return nil
}

func (p *PostgresWalletDbHandler) WalletStatus(id int) (*models.Wallet, error) {
	// Implementar la lógica para obtener el status de una wallet de la base de datos PostgreSQL
	var err error
	var wallet models.Wallet

	stmt, err := db.DbConn.Prepare("SELECT * FROM wallets WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&wallet.ID, &wallet.Balance, &wallet.Currency, &wallet.LogID, &wallet.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNoMatch
	} else if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (p *PostgresWalletDbHandler) UpdateWallet(wallet *models.Wallet) error {
	// Implementar la lógica para actualizar una wallet en la base de datos PostgreSQL

	return nil
}

func (p *PostgresWalletDbHandler) DeleteWallet(id int) error {
	// Implementar la lógica para eliminar una wallet de la base de datos PostgreSQL
	var err error

	stmt, err := db.DbConn.Prepare("DELETE FROM wallets WHERE id = $1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID no encontrado: %d", id)
	}

	return nil
}
