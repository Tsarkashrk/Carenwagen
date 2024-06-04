package domain

type CarRepository interface {
    GetAllCars() ([]*Car, error)
    GetCarByID(id string) (*Car, error)
}
