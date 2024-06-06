package testing

import (
	"bytes"
	"encoding/json"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/data"
	"net/http"
	"testing"
)

func TestEditOrderHandler(t *testing.T) {
	client := &http.Client{}

	updatedCar := data.Car{
		Title:        "Wonderful Dark Green Color",
		Model:        "M5 F90 Competition",
		Make:         "BMW",
		Year:         2021,
		Color:        "dark green",
		Price:        49999990,
		Mileage:      0,
		Description:  "Brand new BMW M5 F90 Competition color dark green without mileage",
		Transmission: "adaptive",
		FuelType:     "petrol",
		ImageURL:     "new_img_url",
	}

	jsonUpdatedCar, _ := json.Marshal(updatedCar)

	req, err := http.NewRequest("PUT", "http://localhost:4000/v1/cars/14", bytes.NewBuffer(jsonUpdatedCar))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
