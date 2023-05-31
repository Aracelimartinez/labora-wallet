package models

import (
	"errors"
	"time"
)

type User struct {
	ID             int       `json:"id"`
	UserName       string    `json:"user_name"`
	DocumentNumber string    `json:"document_number"`
	DocumentType   string    `json:"document_type"`
	Country        string    `json:"country"`
	DateOfBirth    string    `json:"date_of_birth"`
	CreatedAt      time.Time `json:"created_at"`
}

func ValidateUser(user *User) error {
	switch {
	case user.UserName == "":
		return errors.New("el nombre de usuário es obligatório")
	case user.DocumentNumber == "":
		return errors.New("el número de documento es obligatório")
	case user.DocumentType == "":
		return errors.New("el tipo de documento es obligatório")
	case user.Country == "":
		return errors.New("el país es obligatório")
	case user.DateOfBirth == "":
		return errors.New("la fecha de nacimiento es obligatória")
	}
	return nil
}
