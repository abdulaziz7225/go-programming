package main

import (
	"context"
	"log"

	"golang.source-fellows.com/training/carman-grpc/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	client := proto.NewCarServiceClient(con)

	ctx := context.Background()
	autoToCreate := proto.Car {
		Vin: "WVWZZZ1JZ3W386752",
		Color: "babylau",
	}

	_, err = client.CreateCar(ctx, &autoToCreate)
	if err != nil {
		log.Fatalf("Failed to create car: %v", err)
	}

	log.Println("Car successfully sent to the server!")
}