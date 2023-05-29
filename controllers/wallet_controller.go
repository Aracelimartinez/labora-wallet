package controllers

import (
	"encoding/json"
	"labora-wallet/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Función para crear una billetera
func CreateWallet(w http.ResponseWriter, r *http.Request) {

}

// Función para actualizar una billetera
func UpdateWallet(w http.ResponseWriter, r *http.Request) {

}

// Función para excluir una billetera
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

// Función para saber el status de la billetera
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
