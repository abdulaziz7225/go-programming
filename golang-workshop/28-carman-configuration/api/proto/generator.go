package proto

//go:generate protoc --proto_path=. --go_opt=paths=source_relative --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative carman.proto
