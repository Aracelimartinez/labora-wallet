package controllers

import (
	"encoding/json"
	"io/ioutil"
	"labora-wallet/models"
	"labora-wallet/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Function to create a wallet
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	userParam := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userParam)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error al convertir id a entero"))
	}





	// var newWallet models.Wallet

	// err = json.Unmarshal(requestBody, &newUser)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err = services.WS.CreateWallet(newWallet, newLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear la billetera"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Billetera creada con éxito!"))
}

// Function to update a wallet
func UpdateWallet(w http.ResponseWriter, r *http.Request) {

}

// Function to delete a wallet
func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	params := mux.Vars(r)
	idWallet, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error al convertir id a entero"))
		return
	}

	err = services.WS.DeleteWallet(idWallet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al eliminar la billetera"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("La billetera fue eliminada correctamente"))
}

// Function to get the wallet
func WalletStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	params := mux.Vars(r)
	idWallet, err := strconv.Atoi(params["id"])

	wallet, err := services.WS.GetWallet(idWallet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener las informaciones de la billetera"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)

}
