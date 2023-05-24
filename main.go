package main

import (
	// "labora-wallet/db"
	"labora-wallet/controllers"
	// "labora-wallet/services"
	"github.com/gorilla/mux"
)

func main()  {


	router:= mux.NewRouter()

	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/UpdateWallet", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/DeleteWallet", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus", controllers.WalletStatus).Methods("GET")

}
