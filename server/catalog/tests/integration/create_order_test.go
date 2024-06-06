package testing

import (
	"bytes"
	"encoding/json"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/data"
	"net/http"
	"testing"
)

func TestCreateOrderHandler(t *testing.T) {
	client := &http.Client{}

	car := data.Car{
		Title:        "Amazing sport color",
		Model:        "M5 F90",
		Make:         "BMW",
		Year:         2021,
		Color:        "sport blue",
		Price:        50000000,
		Mileage:      0,
		Description:  "Brand new BMW M5 F90 color sport blue without mileage",
		Transmission: "sport",
		FuelType:     "petrol",
		ImageURL:     "something",
	}

	jsonCar, _ := json.Marshal(car)

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/cars", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}
