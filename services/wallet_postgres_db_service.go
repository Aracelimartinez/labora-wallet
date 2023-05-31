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

var errWalletNoMatch = errors.New("billetera no encontrada: Este id no existe")

// Function to create a wallet in PostgreSQL database
func (p *PostgresWalletDbHandler) CreateWallet(user *models.User, log *models.Log) (error) {
	var err error
	var newWallet models.Wallet

	newWallet.Currency = getCurrencyByCountry(user.Country)
	newWallet.LogID = log.ID

	stmt, err := db.DbConn.Prepare("INSERT INTO public.wallets(currency, log_id) VALUES ($1, $2)")
	if err != nil {
		return  err
	}

	defer stmt.Close()

	_, err = stmt.Exec( newWallet.Currency, newWallet.LogID)
	if err != nil {
		return err
	}

	return nil
}

// Function to get the wallet status in PostgreSQL database
func (p *PostgresWalletDbHandler) WalletStatus(id int) (*models.Wallet, error) {
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
		return nil, errWalletNoMatch
	} else if err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Function to update a wallet in PostgreSQL database
func (p *PostgresWalletDbHandler) UpdateWallet(wallet *models.Wallet) error {

	return nil
}

// Function to delete a wallet in PostgreSQL database
func (p *PostgresWalletDbHandler) DeleteWallet(id int) error {
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

func getCurrencyByCountry(country string) string {
	var currency string
	switch country {
	case "BR":
		currency = "R$ - Reales"
	case "PE":
		currency = " S/ - Sol"
	case "CO":
		currency = "$ - Peso colombiano"
	case "CL":
		currency = "$ - Peso chileno"
	case "MX":
		currency = "$ - Peso mexicano"
	case "CR":
		currency = "₡ - Colón costarricense"
		}
	return currency
}
