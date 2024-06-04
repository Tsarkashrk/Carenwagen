package main

import (
	"fmt"

	"github.com/Tsarkashrk/Carenwagen/server/catalog/configs"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/delivery/http"
	"github.com/Tsarkashrk/Carenwagen/server/catalog/internal/usecase"
)

func main() {
	fmt.Println("Hello, World!")
	cfg := configs.LoadConfig()
	carUsecase := usecase.NewCarUsecase(cfg)
	httpHandler := http.NewHandler(carUsecase)
	httpHandler.Run(cfg.Server.Port)
}
