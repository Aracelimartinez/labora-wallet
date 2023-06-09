package main

import (
	"labora-wallet/controllers"
	"labora-wallet/db"
	"log"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := db.EstablishDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	//Wallet endpoints
	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/DeleteWallet", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus/{id}", controllers.WalletStatus).Methods("GET")
	router.HandleFunc("/Wallet/{id}", controllers.GetWalletAndTransactions).Methods("GET")

	//User endpoints
	router.HandleFunc("/CreateUser", controllers.CreateUser).Methods("POST")

	//Transaction endpoints
	router.HandleFunc("/transaction", controllers.CreateTransaction).Methods("POST")

	// Configurate the middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5432"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := corsOptions.Handler(router)

	port := ":8080"
	if err := db.StartServer(port, handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
