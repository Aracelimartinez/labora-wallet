package models

type DbWalletHandler interface {
	CreateWallet(user *User,  log *Log) error
	WalletStatus(id int) (*Wallet, error)
	UpdateWalletBalance(newBalance float64, wallet *Wallet) error
	DeleteWallet(id int) error
	GetWalletAndTransactions(id int) (*WalletDTO, error)
}

type DbUserHandler interface {
	CreateUser(newUser *User) error
	GetUser(id int) (*User, error)
	// UpdateUser(user User) error
	// DeleteUser(id int) error
}

type DbLogHandler interface {
	CreateLog(user *User, canCreate bool) (Log, error)
}

type DbTransactionHandler interface {
	CreateTransaction(newTransaction *Transaction) error
}
