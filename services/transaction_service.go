package services

import (
	"labora-wallet/db"
	"labora-wallet/models"
)

type TransactionService struct {
	DbHandler models.DbTransactionHandler
}

var dbTransactionHandler = &PostgresTransactionDbHandler{Db: db.DbConn}
var TS = &TransactionService{DbHandler: dbTransactionHandler}

func (s *TransactionService) CreateTransaction(newTransaction *models.Transaction) error {
	return s.DbHandler.CreateTransaction(newTransaction)
}
