package models

type DbWalletHandler interface {
	CreateWallet(NewWallet *Wallet,  log *Log) error
	WalletStatus(id int) (*Wallet, error)
	UpdateWallet(wallet *Wallet) error
	DeleteWallet(id int) error
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
