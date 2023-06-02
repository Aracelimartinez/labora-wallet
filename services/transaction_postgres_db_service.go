package services

import (
	"database/sql"
	"errors"
	"fmt"
	"labora-wallet/db"
	"labora-wallet/models"
)

type PostgresTransactionDbHandler struct {
	Db *sql.DB
}

var errInsuficientBalanceFounds = errors.New("no possee suficiente dinero en su cuenta para realizar esta transacción")

// Function to create a transaction in PostgreSQL database
func (p *PostgresTransactionDbHandler) CreateTransaction(newTransaction *models.Transaction) error {
	var err error

	err = newTransaction.ValidateTransactionInfo()
	if err != nil {
		return err
	}

	tx, err := db.DbConn.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar la transacción: %w", err)
	}

	err = executeTransaction(newTransaction)
	if err != nil && err != errInsuficientBalanceFounds {
		return fmt.Errorf("error al ejecutar la transacción: %w", err)
	}

	newTransaction.Approved = true

	stmt, err := p.Db.Prepare("INSERT INTO public.transaction(user_name, document_number, document_type, country, date_of_birth) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newTransaction)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error al hacer commit de la transacción: %w", err)
	}

	if newTransaction.Approved {
		return errInsuficientBalanceFounds
	}

	return nil
}

func executeTransaction(newTransaction *models.Transaction) error {
	WalletMutex.Lock()
	defer WalletMutex.Unlock()

	senderWallet, err := WS.GetWallet(newTransaction.SenderWalletID)
	if err != nil {
		return fmt.Errorf("error al localizar la billetera de origen: %w", err)
	}

	if newTransaction.Type == "Transferencia" {
		receiverWallet, err := WS.GetWallet(newTransaction.ReceiverWalletID)
		if err != nil {
			return fmt.Errorf("error al localizar la billetera que recibirá la transferencia: %w", err)
		}

		err = movement("Retiro", newTransaction.Amount, senderWallet)
		if err != nil {
			return err
		}

		err = movement("Depósito", newTransaction.Amount, receiverWallet)
		if err != nil {
			return err
		}

	} else {
		err = movement(newTransaction.Type, newTransaction.Amount, senderWallet)
		if err != nil {
			return err
		}
	}
	return nil
}

func movement(transactionType string, transactionAmount float64, wallet *models.Wallet) error {
	var err error
	var newBalance float64

	switch transactionType {
	case "Depósito":
		newBalance = wallet.Balance + transactionAmount
	case "Retiro":
		newBalance = wallet.Balance - transactionAmount
		if newBalance < 0.00 {
			return errInsuficientBalanceFounds
		}
	}

	WS.UpdateWalletBalance(newBalance, wallet)
	if err != nil {
		return fmt.Errorf("error al actualizar el balance de la billetera: %w", err)
	}

	return nil
}
