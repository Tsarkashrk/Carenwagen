package usecase

import "github.com/Tsarkashrk/Carenwagen/server/catalog/internal/domain"

type CarUsecase struct {
	carRepository domain.CarRepository
}

func NewCarUsecase(repo domain.CarRepository) *CarUsecase {
	return &CarUsecase{
		carRepository: repo,
	}
}

func (uc *CarUsecase) GetAllCars() ([]*domain.Car, error) {
	return uc.carRepository.GetAllCars()
}

func (uc *CarUsecase) GetCarByID(id string) (*domain.Car, error) {
	return uc.carRepository.GetCarByID(id)
}
