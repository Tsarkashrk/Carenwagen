package testing

import (
	"bytes"
	"encoding/json"
	"github.com/erk1nqz/order-service/internal/data"
	"net/http"
	"testing"
)

func TestCreateOrderHandler(t *testing.T) {
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
}
