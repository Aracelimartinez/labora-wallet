package services

import (
	"labora-wallet/models"
)

func TryToCreateWallet(user models.User) error {
	canCreate, err := CheckIfCanCreateWallet(user)
	if err != nil {

		return err
	}
	if canCreate {
		//Te creas log success

		return nil
	}
	//Te creas log rejected
	return nil
}

func CheckIfCanCreateWallet(user models.User) (bool, error) {
	//Toda tu logica de llamar a la API de truora

	//POST a la Api de Checks de Truora --> Llevarlo a un func
	//GET a la Api de Checks de Truora --> Llevarlo a un func

	return true, nil
}
