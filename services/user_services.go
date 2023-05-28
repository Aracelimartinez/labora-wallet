package services

import (
	"labora-wallet/db"
	"labora-wallet/models"
)

type UserService struct {
	DbHandler models.DbUserHandler
}

var dbUserHandler = &PostgresUserDbHandler{Db: db.DbConn}
var US = &UserService{DbHandler: dbUserHandler}

func (s *UserService) CreateUser(user models.User) error {
	return s.DbHandler.CreateUser(user)
}

// func (s *UserService) GetUser(id int) (models.User, error) {
// 	return s.DbHandler.GetUser(id)
// }

// func (s *UserService) UpdateUser(user models.User) error {
// 	return s.DbHandler.UpdateUser(user)
// }

// func (s *UserService) DeleteUser(id int) error {
// 	return s.DbHandler.DeleteUser(id)
// }
