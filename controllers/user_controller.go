package controllers

import (
	"encoding/json"
	"io/ioutil"
	"labora-wallet/models"
	"labora-wallet/services"
	"log"
	"net/http"
)

// Función para crear una billetera
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var newUser models.User

	err = json.Unmarshal(requestBody, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.US.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear el usuário"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário creado con éxito!"))
}

// Función para actualizar una billetera
// func UpdateUser(w http.ResponseWriter, r *http.Request) {

// }

// Función para excluir una billetera
// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// }

// Función para saber el status de la billetera
// func GetUser(w http.ResponseWriter, r *http.Request) {

// }
