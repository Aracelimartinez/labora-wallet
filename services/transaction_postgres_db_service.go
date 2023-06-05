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

	err = executeTransaction(newTransaction)
	if err != nil && err != errInsuficientBalanceFounds {

		return fmt.Errorf("error al ejecutar la transacción: %w", err)
	}

	if err == errInsuficientBalanceFounds {
		newTransaction.Approved = false
	} else {
		newTransaction.Approved = true
	}

	stmt, err := db.DbConn.Prepare("INSERT INTO public.transactions(senders_wallet_id, receiver_wallet_id, type, amount, date, approved) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {

		return err
	}

	_, err = stmt.Exec(newTransaction.SenderWalletID, newTransaction.ReceiverWalletID, newTransaction.Type, newTransaction.Amount, newTransaction.Date, newTransaction.Approved)
	if err != nil {

		return err
	}

	defer stmt.Close()

	if !newTransaction.Approved {
		return errInsuficientBalanceFounds
	}

	return nil
}

func executeTransaction( newTransaction *models.Transaction) error {
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

		err = movementInTransaction("Retiro", newTransaction.Amount, senderWallet)
		if err != nil {
			return err
		}

		err = movementInTransaction("Depósito", newTransaction.Amount, receiverWallet)
		if err != nil {
			return err
		}

	} else {
		err = movementInTransaction( newTransaction.Type, newTransaction.Amount, senderWallet)
		if err != nil {
			return err
		}
	}
	return nil
}

func movementInTransaction(transactionType string, transactionAmount float64, wallet *models.Wallet) error {
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

	WS.UpdateWalletBalance( newBalance, wallet)
	if err != nil {
		return fmt.Errorf("error al actualizar el balance de la billetera: %w", err)
	}

	return nil
}

// Function to get transactions by wallet ID
func (p *PostgresWalletDbHandler) GetTransactionsByWalletID(walletID int) ([]models.Transaction, error) {
	var err error
	var transactions []models.Transaction

	stmt, err := db.DbConn.Prepare("SELECT * FROM transactions WHERE sender_wallet_id = $1 OR receiver_wallet_id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(walletID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction.ID, &transaction.SenderWalletID, &transaction.ReceiverWalletID, &transaction.Type, &transaction.Amount, &transaction.Date, &transaction.Approved, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
