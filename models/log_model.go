package models

import "time"

type Log struct {
	ID        int       `json:"id"`
	Approved  bool      `json:"approved"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
}
