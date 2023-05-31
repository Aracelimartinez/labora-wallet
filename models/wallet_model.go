package models

type Wallet struct {
	ID        int     `json:"id"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	LogID     int     `json:"log_id"`
	CreatedAt string  `json:"created_at"`
}


