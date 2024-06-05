package testing

import (
	"bytes"
	"encoding/json"
	"github.com/erk1nqz/order-service/internal/data"
	"net/http"
	"testing"
)

func TestEditOrderHandler(t *testing.T) {
	client := &http.Client{}

	order := data.Order{
		CustomerID: 1,
		CarModel:   "BMW M5",
		Color:      "Red",
		Load:       "Basic",
	}

	jsonOrder, _ := json.Marshal(order)

	req, err := http.NewRequest("PUT", "http://localhost:4000/v1/order/4", bytes.NewBuffer(jsonOrder))
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
