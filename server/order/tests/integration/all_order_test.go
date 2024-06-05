package testing

import (
	"bytes"
	"encoding/json"
	"github.com/erk1nqz/order-service/internal/data"
	"net/http"
	"testing"
)

func TestCreateEditDeleteOrderHandler(t *testing.T) {
	client := &http.Client{}

	order := data.Order{
		CustomerID: 1,
		CarModel:   "BMW M5",
		Color:      "Red",
		Load:       "Competition",
	}

	jsonOrder, _ := json.Marshal(order)

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/order", bytes.NewBuffer(jsonOrder))
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

	updatedOrder := data.Order{
		CustomerID: 1,
		CarModel:   "BMW M5",
		Color:      "Blue",
		Load:       "Basic",
	}

	jsonUpdatedOrder, _ := json.Marshal(updatedOrder)

	req1, err := http.NewRequest("PUT", "http://localhost:4000/v1/order/5", bytes.NewBuffer(jsonUpdatedOrder))
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

	req2, err := http.NewRequest("DELETE", "http://localhost:4000/v1/order/5", nil)
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
