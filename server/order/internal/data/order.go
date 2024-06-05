package data

import "time"

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	CarID      int64     `json:"car_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
