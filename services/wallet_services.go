package services

import (
	"labora-wallet/models"
)

type WalletService struct {
	DbHandler DBWalletHandler
}

func (s *WalletService) CreateWallet(wallet models.Wallet) error {
	return s.DbHandler.CreateWallet(wallet)
}

func (s *WalletService) GetWallet(id int) (models.Wallet, error) {
	return s.DbHandler.WalletStatus(id)
}

func (s *WalletService) UpdateWallet(wallet models.Wallet) error {
	return s.DbHandler.UpdateWallet(wallet)
}

func (s *WalletService) DeleteWallet(id int) error {
	return s.DbHandler.DeleteWallet(id)
}
