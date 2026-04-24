package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	carman "golang.source-fellows.com/training/carman/v7/internal"
	"golang.source-fellows.com/training/carman/v7/internal/model"
)

type CarRepository struct {
	cars []model.Car
}

func (cr *CarRepository) AddCar(ctx context.Context, car model.Car) error {
	logrus.WithContext(ctx).Info("Adding new car to memory database")

	logger := carman.Logger(ctx)
	logger.Info("AddCar called")

	cr.cars = append(cr.cars, car)
	return nil
}

func (cr *CarRepository) GetAllCars(ctx context.Context) ([]model.Car, error) {
	logrus.WithContext(ctx).Info("Fetching all cars from memory database")

	logger := carman.Logger(ctx)
	logger.Info("GetAllCars called")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("wait complete")
	case <-ctx.Done():
		logrus.WithContext(ctx).Warn("Request was cancelled by the client")
		return nil, ctx.Err()
	}
	return cr.cars, nil
}
