.PHONY: grpc-add rest-get

# Run this by typing: make grpc-add
grpc-add:
	grpcurl -plaintext -import-path ./api/proto -proto carman.proto \
	-d '{"plate": "GRPC-MAKE", "speed_in_km_h": 220, "engine_started": true, "transmission": {"rotations": 4000, "engaged_gear": 5}}' \
	localhost:9090 proto.CarService/AddAudi

# Run this by typing: make rest-get
rest-get:
	curl -X GET http://localhost:8080/api/cars -u "foo:bar"