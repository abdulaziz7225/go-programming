package internal

import "golang.source-fellows.com/training/carman/v2/internal/model"

type CarRepository interface {
	AddCar(car model.Car) error
	GetAllCars() ([]model.Car, error)
}
