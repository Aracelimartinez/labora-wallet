package services

import (
	"database/sql"
	"errors"
	"fmt"
	"labora-wallet/models"
	"sync"
)

type PostgresWalletDbHandler struct {
	Db *sql.DB
}

var errWalletNoMatch = errors.New("billetera no encontrada: Este id no existe")
var WalletMutex sync.Mutex

// Function to create a wallet in PostgreSQL database
func (p *PostgresWalletDbHandler) CreateWallet(user *models.User, log *models.Log) error {
	var err error
	var newWallet models.Wallet
	WalletMutex.Lock()
	newWallet.AccountNumber = models.GenerateUniqueAccountNumber()
	newWallet.Currency = models.SetCurrencyByCountry(user.Country)
	newWallet.LogID = log.ID

	stmt, err := p.Db.Prepare("INSERT INTO public.wallets(account_number, currency, log_id) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newWallet.AccountNumber, newWallet.Currency, newWallet.LogID)
	if err != nil {
		return err
	}

	WalletMutex.Unlock()

	return nil
}

// Function to get the wallet status in PostgreSQL database
func (p *PostgresWalletDbHandler) WalletStatus(id int) (*models.Wallet, error) {
	var err error
	var wallet models.Wallet

	stmt, err := p.Db.Prepare("SELECT * FROM wallets WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&wallet.ID, &wallet.AccountNumber, &wallet.Balance, &wallet.Currency, &wallet.LogID, &wallet.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errWalletNoMatch
	} else if err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Function to update a wallet balance in PostgreSQL database
func (p *PostgresWalletDbHandler) UpdateWalletBalance(newBalance float64, wallet *models.Wallet) error {
	var err error

	stmt, err := p.Db.Prepare("UPDATE wallets	SET balance = $1 WHERE id = $2")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newBalance, wallet.ID)
	if err != nil {
		return err
	}

	return nil
}

// Function to delete a wallet in PostgreSQL database
func (p *PostgresWalletDbHandler) DeleteWallet(id int) error {
	var err error

	stmt, err := p.Db.Prepare("DELETE FROM wallets WHERE id = $1")
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
