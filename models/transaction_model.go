package models

import (
	"errors"
	"time"
)

type Transaction struct {
	ID               int       `json:"id"`
	SenderWalletID   int       `json:"sender_wallet_id"`
	ReceiverWalletID int       `json:"receiver_wallet_id"`
	Type             string    `json:"type"`
	Amount           float64   `json:"amount"`
	Date             string    `json:"date"`
	Approved         bool      `json:"approved"`
	CreatedAt        time.Time `json:"created_at"`
}

func (transaction *Transaction) ValidateTransactionInfo() error {
	var err error

	err = transaction.ValidateTransactionType()
	if err != nil {
		return err
	}

	err = transaction.ValidateTransactionDate()
	if err != nil {
		return err
	}

	err = transaction.validateTransferType()
	if err != nil {
		return err
	}

	return nil
}

func (transaction *Transaction) ValidateTransactionType() error {

	if transaction.Type != "Retiro" && transaction.Type != "Transferencia" && transaction.Type != "Depósito" {
		err := errors.New("por favor inserte un tipo de transaction válido: Retiro, Depósito o Transferencia ")
		return err
	}
	return nil
}

func (transaction *Transaction) ValidateTransactionDate() error {
	var err error
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", transaction.Date)
	if err != nil {
		return err
	}
	now := time.Now()

	if t.Before(now) {
		err = errors.New("fecha de transacción inválida")
		return err
	}
	return nil
}

func (transaction *Transaction) validateTransferType() error {
	var err error
	if transaction.Type == "Transferencia" && transaction.ReceiverWalletID == 0 {
		err = errors.New("la billetera que recibirá la transferencia es obligatoria")
		return err
	}
	return nil
}
