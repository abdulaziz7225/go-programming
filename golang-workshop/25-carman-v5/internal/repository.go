package internal

import (
	"context"

	"golang.source-fellows.com/training/carman/v5/internal/model"
)

type CarRepository interface {
	AddCar(ctx context.Context, car model.Car) error
	GetAllCars(ctx context.Context) ([]model.Car, error)
}
