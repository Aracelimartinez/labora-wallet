package models

import (
	"math/rand"
	"time"
)

type Wallet struct {
	ID            int     `json:"id"`
	AccountNumber int     `json:"account_number"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	LogID         int     `json:"log_id"`
	CreatedAt     string  `json:"created_at"`
}
type WalletDTO struct {
	Wallet       *Wallet        `json:"id"`
	Transactions []Transaction `json:"transactions"`
}

func SetCurrencyByCountry(country string) string {
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

func GenerateUniqueAccountNumber() int {
	var accountNumber int
	now := time.Now()

	rand.Seed(time.Now().UnixNano())

	accountNumber = rand.Intn(100000-1000+1) + 1000
	lastFourDigits := now.Minute()*100 + now.Second()
	accountNumberWithTime := accountNumber*10000 + lastFourDigits
	return accountNumberWithTime
}
