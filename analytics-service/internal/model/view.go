package model

import "time"

type View struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Timestamp time.Time `json:"timestamp"`
}
