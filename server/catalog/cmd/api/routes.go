package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/cars", app.createCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars/:id", app.getCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars", app.getAllCarsHandler)
	router.HandlerFunc(http.MethodPut, "/v1/cars/:id", app.editCarHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/cars/:id", app.deleteCarHandler)

	return app.recoverPanic(app.rateLimit(enableCORS(router)))
}
