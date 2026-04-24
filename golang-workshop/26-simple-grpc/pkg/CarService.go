package pkg

import (
	"context"
	"fmt"

	"golang.source-fellows.com/training/carman-grpc/pkg/proto"
)

type CarServiceImpl struct {
	proto.UnimplementedCarServiceServer
}

func (s *CarServiceImpl) CreateCar(ctx context.Context, car *proto.Car) (*proto.Void, error) {
	// We use the generated GetVin() and GetColor() getter methods
	fmt.Printf("Successfully received new car - VIN: %s | Color: %s\n", car.GetVin(), car.GetColor())

	// Return the empty Void message
	return &proto.Void{}, nil
}
