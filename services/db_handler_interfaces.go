package services

import(
	"labora-wallet/models"
)

type DBWalletHandler interface {
	CreateWallet(wallet models.Wallet) error
	WalletStatus(id int) (models.Wallet, error)
	UpdateWallet(wallet models.Wallet) error
	DeleteWallet(id int) error
}
