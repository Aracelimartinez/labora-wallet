package main

import (
	"labora-wallet/db"
	"labora-wallet/controllers"
	"log"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func main()  {

	router:= mux.NewRouter()

	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/UpdateWallet", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/DeleteWallet", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus", controllers.WalletStatus).Methods("GET")

	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5432"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := corsOptions.Handler(router)

	port := ":8000"
	if err := db.StartServer(port, handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
