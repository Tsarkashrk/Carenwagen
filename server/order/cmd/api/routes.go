package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodPost, "/v1/order", app.createOrderHandler)
	router.HandlerFunc(http.MethodGet, "/v1/order/:id", app.getOrderHandler)
	router.HandlerFunc(http.MethodGet, "/v1/order", app.getAllOrdersHandler)
	router.HandlerFunc(http.MethodPut, "/v1/order/:id", app.editOrderHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/order/:id", app.deleteOrderHandler)
	return app.recoverPanic(app.rateLimit(enableCORS(router)))
}
