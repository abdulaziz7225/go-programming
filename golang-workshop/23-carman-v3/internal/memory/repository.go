package memory

import "golang.source-fellows.com/training/carman/v3/internal/model"

type CarRepository struct {
	cars []model.Car
}

func (cr *CarRepository) AddCar(car model.Car) error {
	cr.cars = append(cr.cars, car)
	return nil
}

func (cr *CarRepository) GetAllCars() ([]model.Car, error) {
	return cr.cars, nil
}