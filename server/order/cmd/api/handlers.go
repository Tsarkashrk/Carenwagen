package main

import (
	"github.com/erk1nqz/order-service/internal/data"
	"net/http"
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
	order.ID = int(id)

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
