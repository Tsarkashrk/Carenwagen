package main

import (
	"github.com/Tsarkashrk/Carenwagen/server/order/internal/data"
	"net/http"
	"strconv"
)

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order data.Order
	if err := app.readJSON(w, r, &order); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.db.Insert(&order); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"order": order}, nil)
}

func (app *application) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	order, err := app.db.Retrieve(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
}

func (app *application) getAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	page := 1
	pageSize := 10
	sort := "ASC"
	carID := 0

	query := r.URL.Query()

	if p, ok := query["page"]; ok {
		if parsedPage, err := strconv.Atoi(p[0]); err == nil {
			page = parsedPage
		}
	}
	if ps, ok := query["pageSize"]; ok {
		if parsedPageSize, err := strconv.Atoi(ps[0]); err == nil {
			pageSize = parsedPageSize
		}
	}
	if s, ok := query["sort"]; ok {
		if s[0] == "DESC" {
			sort = "DESC"
		}
	}
	if c, ok := query["car_id"]; ok {
		if parsedCarID, err := strconv.Atoi(c[0]); err == nil {
			carID = parsedCarID
		}
	}

	orders, err := app.db.RetrieveAll(page, pageSize, sort, carID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"orders": orders}, nil)
}

func (app *application) editOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var order data.Order
	if err := app.readJSON(w, r, &order); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	order.ID = int64(id)

	if err := app.db.Update(&order); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.db.Delete(int(id)); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Order deleted successfully"}, nil)
}
