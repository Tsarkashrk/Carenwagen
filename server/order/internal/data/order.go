package data

import "time"

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	CarID      int       `json:"car_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
