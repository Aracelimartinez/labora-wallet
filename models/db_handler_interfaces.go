package models

type DbWalletHandler interface {
	CreateWallet(wallet Wallet) error
	WalletStatus(id int) (Wallet, error)
	UpdateWallet(wallet Wallet) error
	DeleteWallet(id int) error
}

type DbUserHandler interface {
	CreateUser(user User) error
	// GetUser(id int) (User, error)
	// UpdateUser(user User) error
	// DeleteUser(id int) error
}
