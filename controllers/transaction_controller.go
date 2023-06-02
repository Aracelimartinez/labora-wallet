package controllers

import (
	"encoding/json"
	"io/ioutil"
	"labora-wallet/models"
	"labora-wallet/services"
	"log"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var newTransaction models.Transaction

	err = json.Unmarshal(requestBody, &newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.TS.CreateTransaction(&newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al realizar la transacción"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transacción finalizada con éxito!"))
}
