package main

import (
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/data"
	"net/http"
	"strconv"
)

func (app *application) createCarHandler(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	if err := app.readJSON(w, r, &car); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.db.Insert(&car); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"car": car}, nil)
}

func (app *application) getCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	car, err := app.db.Retrieve(int(id))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
}

func (app *application) getAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	page := 1
	pageSize := 10
	sort := "ASC"
	model := ""

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
	if m, ok := query["model"]; ok {
		model = m[0]
	}

	cars, err := app.db.RetrieveAll(page, pageSize, sort, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"cars": cars}, nil)
}

func (app *application) editCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var car data.Car
	if err := app.readJSON(w, r, &car); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	car.ID = int64(id)

	if err := app.db.Update(&car); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
}

func (app *application) deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.db.Delete(int(id)); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Car deleted successfully"}, nil)
}
