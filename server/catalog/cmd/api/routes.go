package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/cars", app.addCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars/:id", app.showCarHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cars", app.showAllCarHandler)
	router.HandlerFunc(http.MethodPut, "/v1/cars/:id", app.updateCarHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/cars/:id", app.deleteCarHandler)
	return router
}
