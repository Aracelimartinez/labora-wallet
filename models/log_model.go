package models

type Log struct {
	ID				int			`json:"id"`
	Aproved		bool		`json:"aproved"`
	CreatedAt string	`json:"created_at"`
	UserID 		int 		`json:"user_id"`
}
