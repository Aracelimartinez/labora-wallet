package models

type User struct {
	ID							int			`json:"id"`
	UserName				string	`json:"user_name"`
	DocumentNumber	string	`json:"document_number"`
	DocumentType		string	`json:"document_type"`
	Country 				string	`json:"country"`
}
