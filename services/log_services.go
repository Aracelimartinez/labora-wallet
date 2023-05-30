package services

import (
	"labora-wallet/db"
	"labora-wallet/models"
)

type LogService struct {
	DbHandler models.DbLogHandler
}

var dbLogHandler = &PostgresLogDbHandler{Db: db.DbConn}
var LS = &LogService{DbHandler: dbLogHandler}

func (s *LogService) CreateLog(user *models.User, canCreate bool ) (models.Log, error) {
	return s.DbHandler.CreateLog(user, canCreate)
}
