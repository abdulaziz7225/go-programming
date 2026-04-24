package main

import (
	"log"
	"net"

	"golang.source-fellows.com/training/carman-grpc/pkg"
	"golang.source-fellows.com/training/carman-grpc/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()

	proto.RegisterCarServiceServer(srv, &pkg.CarServiceImpl{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on :8080")
	log.Fatal(srv.Serve(listener))
}
