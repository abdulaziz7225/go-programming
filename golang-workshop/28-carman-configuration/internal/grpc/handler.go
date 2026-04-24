package grpc

import (
	"context"

	"github.com/sirupsen/logrus"
	carman "golang.source-fellows.com/training/carman/v7/internal"
	"golang.source-fellows.com/training/carman/v7/internal/model"
	"golang.source-fellows.com/training/carman/v7/api/proto"
)

// GrpcHandler implements the generated Protobuf interface
type GrpcHandler struct {
	proto.UnimplementedCarServiceServer
	Repo carman.CarRepository // Reusing your existing database!
}

func (h *GrpcHandler) AddAudi(ctx context.Context, req *proto.Audi) (*proto.Void, error) {
	logrus.WithContext(ctx).Info("gRPC: Received request to add Audi")

	// Domain Mapping: Convert the Protobuf request into your internal Go model
	newAudi := model.Audi{
		Plate:         model.LicensePlate(req.GetPlate()),
		SpeedInKmH:    int(req.GetSpeedInKmH()),
		EngineStarted: req.GetEngineStarted(),
	}

	if req.Transmission != nil {
		newAudi.Transmission = &model.Transmission{
			Rotations:   int(req.Transmission.GetRotations()),
			EngagedGear: int(req.Transmission.GetEngagedGear()),
		}
	}

	// Save it using your existing repository
	if err := h.Repo.AddCar(ctx, &newAudi); err != nil {
		return nil, err
	}

	return &proto.Void{}, nil
}