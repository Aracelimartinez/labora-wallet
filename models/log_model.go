package models

type Log struct {
	ID        int    `json:"id"`
	Approved  bool   `json:"approved"`
	CreatedAt string `json:"created_at"`
	UserID    int    `json:"user_id"`
}
