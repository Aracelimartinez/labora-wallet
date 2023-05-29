package services

import (
	"labora-wallet/db"
	"labora-wallet/models"
)

type WalletService struct {
	DbHandler models.DbWalletHandler
}

var dbWalletHandler = &PostgresWalletDbHandler{Db: db.DbConn}
var WS = &WalletService{DbHandler: dbWalletHandler}

func (s *WalletService) CreateWallet(newWallet *models.Wallet,  log *models.Log) error {
	return s.DbHandler.CreateWallet(newWallet, log)
}

func (s *WalletService) GetWallet(id int) (*models.Wallet, error) {
	return s.DbHandler.WalletStatus(id)
}

func (s *WalletService) UpdateWallet(wallet *models.Wallet) error {
	return s.DbHandler.UpdateWallet(wallet)
}

func (s *WalletService) DeleteWallet(id int) error {
	return s.DbHandler.DeleteWallet(id)
}
