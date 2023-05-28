package models

import "errors"

type User struct {
	ID             int    `json:"id"`
	UserName       string `json:"user_name"`
	DocumentNumber string `json:"document_number"`
	DocumentType   string `json:"document_type"`
	Country        string `json:"country"`
}

func ValidateUser(user *User) error {
	switch {
	case user.UserName == "":
		return errors.New("El nombre de usuário es obligatório")
	case user.DocumentNumber == "":
		return errors.New("El número de documento es obligatório")
	case user.DocumentType == "":
		return errors.New("El tipo de documento es obligatório")
	case user.Country == "":
		return errors.New("El país es obligatório")
	}

	return nil
}
