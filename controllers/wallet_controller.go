package controllers

import (
	"net/http"
	"labora-wallet/db"
	"labora-wallet/services"
)

var dbHandler =  &services.PostgresWalletDBHandler{Db: db.Db}
var walletService = &services.WalletService{DbHandler: dbHandler}

// Función para crear una billetera
func CreateWallet(w http.ResponseWriter, r *http.Request)  {

}

// Función para actualizar una billetera
func UpdateWallet(w http.ResponseWriter, r *http.Request)  {

}

// Función para excluir una billetera
func DeleteWallet(w http.ResponseWriter, r *http.Request)  {

}

//Función para saber el status de la billetera
func WalletStatus(w http.ResponseWriter, r *http.Request)  {

}
