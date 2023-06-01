package models

import "time"

type Transaction struct {
	ID               int       `json:"id"`
	SenderWalletID   int       `json:"sender_wallet_id"`
	ReceiverWalletID int       `json:"receiver_wallet_id"`
	Ammount          int       `json:"ammount"`
	Date             string    `json:"date"`
	CreatedAt        time.Time `json:"created_at"`
}
