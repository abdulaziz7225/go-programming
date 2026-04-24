package main

import (
	"log"
	"net"

	"github.com/kelseyhightower/envconfig"
	"golang.source-fellows.com/training/carman/v7/api/proto"
	mygrpc "golang.source-fellows.com/training/carman/v7/internal/grpc"
	myhttp "golang.source-fellows.com/training/carman/v7/internal/http"
	"golang.source-fellows.com/training/carman/v7/internal/memory"
	"google.golang.org/grpc"
)

func main() {
	// 1. Initialize a single shared database
	repo := &memory.CarRepository{}

	// 2. Start the gRPC Server in a background Goroutine
	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register our handler, passing in the shared repo
		handler := &mygrpc.GrpcHandler{Repo: repo}
		proto.RegisterCarServiceServer(grpcServer, handler)

		log.Println("gRPC Server listening on :9090")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	config := &myhttp.Config{}
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Start the Gin REST Server on the main thread
	log.Printf("REST Server listening on %s\n", config.Binding)
	if err := myhttp.StartServer(*config, repo); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
