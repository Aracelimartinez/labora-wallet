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
