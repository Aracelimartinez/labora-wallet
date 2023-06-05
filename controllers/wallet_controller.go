package controllers

import (
	"encoding/json"
	"labora-wallet/services"
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error al convertir id a entero"))
	}

	user, err := services.US.GetUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener el usuário"))
	}

	logCreated, err := services.TryToCreateWallet(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al validar la creación de la billetera"))
		return
	}

	if !logCreated.Approved {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("La creación de su billetera no fue aprovada"))
	}

	err = services.WS.CreateWallet(user, &logCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear la billetera"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Billetera creada con éxito!"))
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error al convertir id a entero"))
		return
	}

	wallet, err := services.WS.GetWallet(idWallet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener las informaciones de la billetera"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)
}

func GetWalletAndTransactions(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var err error
	params := mux.Vars(r)
	idWallet, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error al convertir id a entero"))
		return
	}

	walletInfo, err := services.WS.GetWalletAndTransactions(idWallet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener las informaciones de la billetera"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(walletInfo)

}
