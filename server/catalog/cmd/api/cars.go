package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/data"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/validator"
)

func (app *application) addCarHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title        string  `json:"title"`
		Model        string  `json:"model"`
		Make         string  `json:"make"`
		Year         int32   `json:"year"`
		Color        string  `json:"color"`
		Price        float64 `json:"price"`
		Mileage      int32   `json:"mileage"`
		Description  string  `json:"description"`
		Transmission string  `json:"transmission"`
		FuelType     string  `json:"fuel_type"`
		ImageURL     string  `json:"image_url"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	car := &data.Car{
		Title:        input.Title,
		Model:        input.Model,
		Make:         input.Make,
		Year:         input.Year,
		Color:        input.Color,
		Price:        input.Price,
		Mileage:      input.Mileage,
		Description:  input.Description,
		Transmission: input.Transmission,
		FuelType:     input.FuelType,
		ImageURL:     input.ImageURL,
	}

	v := validator.New()

	if data.ValidateCar(v, car); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	car := data.Car{
		ID:           id,
		CreatedAt:    time.Now(),
		Title:        "Lamborghini Aventador 2022",
		Make:         "Lamborghini",
		Model:        "Aventador",
		Year:         2022,
		Price:        20000.00,
		Color:        "Red",
		Mileage:      15000,
		Description:  "A well-maintained car with excellent fuel efficiency.",
		Transmission: "Automatic",
		FuelType:     "Gasoline",
		ImageURL:     "https://i.ytimg.com/vi/M-Bvors61aE/maxresdefault.jpg",
		Version:      1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
