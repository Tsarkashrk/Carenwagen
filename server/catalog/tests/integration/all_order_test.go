package testing

import (
	"bytes"
	"encoding/json"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/data"
	"net/http"
	"testing"
)

func TestCreateEditDeleteOrderHandler(t *testing.T) {
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

	req1, err := http.NewRequest("PUT", "http://localhost:4000/v1/cars/13", bytes.NewBuffer(jsonUpdatedCar))
	if err != nil {
		t.Fatal(err)
	}

	resp1, err := client.Do(req1)
	if err != nil {
		t.Fatal(err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	req2, err := http.NewRequest("DELETE", "http://localhost:4000/v1/cars/13", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
