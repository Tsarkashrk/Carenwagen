package http

import (
    "github.com/Tsarkashrk/Carenwagen/server/catalog/internal/usecase"
    "net/http"
)

type Handler struct {
    carUsecase *usecase.CarUsecase
}

func NewHandler(usecase *usecase.CarUsecase) *Handler {
    return &Handler{
        carUsecase: usecase,
    }
}

func (h *Handler) GetAllCars(w http.ResponseWriter, r *http.Request) {
    // Implement logic to get all cars from usecase
}

func (h *Handler) GetCarByID(w http.ResponseWriter, r *http.Request) {
    // Implement logic to get a car by ID from usecase
}

func (h *Handler) Run(port string) {
    http.HandleFunc("/cars", h.GetAllCars)
    http.HandleFunc("/cars/{id}", h.GetCarByID)
    http.ListenAndServe(":"+port, nil)
}
