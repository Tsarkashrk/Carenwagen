package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Insert(order *Order) error {
	_, err := m.DB.Exec("INSERT INTO orders (customer_id, car_id, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		order.CustomerID, order.CarID, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create an order: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (*Order, error) {
	var order Order
	row := m.DB.QueryRow("SELECT * FROM orders WHERE id = $1", id)
	err := row.Scan(&order.ID, &order.CustomerID, &order.CarID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve order: %w", err)
	}
	return &order, nil
}

func (m *DBModel) Update(order *Order) error {
	_, err := m.DB.Exec("UPDATE orders SET created_at = $1, updated_at = $2, customer_id = $3, car_id = $4 WHERE id = $5",
		order.CreatedAt, order.UpdatedAt, order.CustomerID, order.CarID, order.ID)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}
