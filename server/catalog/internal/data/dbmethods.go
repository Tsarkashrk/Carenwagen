package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Insert(car *Car) error {
	query := `
		INSERT INTO cars (created_at, title, model, make, year, color, price, mileage, description, transmission, fuel_type, image_url, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := m.DB.Exec(query, car.CreatedAt, car.Title, car.Model, car.Make, car.Year, car.Color, car.Price, car.Mileage, car.Description, car.Transmission, car.FuelType, car.ImageURL, car.Version)
	if err != nil {
		return fmt.Errorf("failed to create a car: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (*Car, error) {
	var car Car
	query := `
		SELECT id, created_at, title, model, make, year, color, price, mileage, description, transmission, fuel_type, image_url, version
		FROM cars
		WHERE id = $1
	`
	row := m.DB.QueryRow(query, id)
	err := row.Scan(&car.ID, &car.CreatedAt, &car.Title, &car.Model, &car.Make, &car.Year, &car.Color, &car.Price, &car.Mileage, &car.Description, &car.Transmission, &car.FuelType, &car.ImageURL, &car.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve car: %w", err)
	}
	return &car, nil
}

func (m *DBModel) RetrieveAll(page, pageSize int, sort string, model string) ([]*Car, error) {
	var cars []*Car

	query := `
		SELECT id, created_at, title, model, make, year, color, price, mileage, description, transmission, fuel_type, image_url, version
		FROM cars
		WHERE ($1 = '' OR model = $1)
		ORDER BY created_at ` + sort + `
		LIMIT $2 OFFSET $3
	`

	rows, err := m.DB.Query(query, model, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cars: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.CreatedAt, &car.Title, &car.Model, &car.Make, &car.Year, &car.Color, &car.Price, &car.Mileage, &car.Description, &car.Transmission, &car.FuelType, &car.ImageURL, &car.Version)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car: %w", err)
		}
		cars = append(cars, &car)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return cars, nil
}

func (m *DBModel) Update(car *Car) error {
	query := `
		UPDATE cars
		SET created_at = $1, title = $2, model = $3, make = $4, year = $5, color = $6, price = $7, mileage = $8, description = $9, transmission = $10, fuel_type = $11, image_url = $12, version = $13
		WHERE id = $14
	`
	_, err := m.DB.Exec(query, car.CreatedAt, car.Title, car.Model, car.Make, car.Year, car.Color, car.Price, car.Mileage, car.Description, car.Transmission, car.FuelType, car.ImageURL, car.Version, car.ID)
	if err != nil {
		return fmt.Errorf("failed to update car: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	query := `
		DELETE FROM cars WHERE id = $1
	`
	_, err := m.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete car: %w", err)
	}
	return nil
}
