package data

import (
	"time"

	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/validator"
)

type Car struct {
	ID           int64
	CreatedAt    time.Time
	Title        string
	Model        string
	Make         string
	Year         int32
	Color        string
	Price        float64
	Mileage      int32
	Description  string
	Transmission string
	FuelType     string
	ImageURL     string
	Version      int32
}

func ValidateCar(v *validator.Validator, car *Car) {
	v.Check(car.Title != "", "title", "must be provided")
	v.Check(len(car.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(car.Model != "", "model", "must be provided")
	v.Check(len(car.Model) <= 100, "model", "must not be more than 100 bytes long")

	v.Check(car.Make != "", "make", "must be provided")
	v.Check(len(car.Make) <= 100, "make", "must not be more than 100 bytes long")

	v.Check(car.Year != 0, "year", "must be provided")
	v.Check(car.Year >= 1888, "year", "must be greater than 1888")
	v.Check(car.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(car.Color != "", "color", "must be provided")
	v.Check(len(car.Color) <= 50, "color", "must not be more than 50 bytes long")

	v.Check(car.Price >= 0, "price", "must be a positive value")
	v.Check(car.Mileage >= 0, "mileage", "must be a positive value")

	v.Check(car.Description != "", "description", "must be provided")
	v.Check(len(car.Description) <= 1000, "description", "must not be more than 1000 bytes long")

	v.Check(car.Transmission != "", "transmission", "must be provided")
	v.Check(len(car.Transmission) <= 50, "transmission", "must not be more than 50 bytes long")

	v.Check(car.FuelType != "", "fuel_type", "must be provided")
	v.Check(len(car.FuelType) <= 50, "fuel_type", "must not be more than 50 bytes long")

	v.Check(car.ImageURL != "", "image_url", "must be provided")
}
