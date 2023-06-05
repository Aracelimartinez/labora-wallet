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

func (s *WalletService) CreateWallet(user *models.User, log *models.Log) error {
	return s.DbHandler.CreateWallet(user, log)
}

func (s *WalletService) GetWallet(id int) (*models.Wallet, error) {
	return s.DbHandler.WalletStatus(id)
}

func (s *WalletService) UpdateWalletBalance(newBalance float64, wallet *models.Wallet) error {
	return s.DbHandler.UpdateWalletBalance(newBalance, wallet)
}

func (s *WalletService) DeleteWallet(id int) error {
	return s.DbHandler.DeleteWallet(id)
}

func (s *WalletService) GetWalletAndTransactions(id int) (models.WalletDTO, error) {
	return s.DbHandler.GetWalletAndTransactions(id)
}
