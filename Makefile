

build-client: build-grpc
	@docker build -f ./client.Dockerfile -t port-service-client ./src

build-service: build-grpc
	@docker build -f ./service.Dockerfile -t port-service ./src

build-grpc:
	@protoc --proto_path=./src/grpc/proto ./src/grpc/proto/*.proto --go_out=./src --go-grpc_out=./src

run-client:
	@docker run --network=host port-service-client

run-service:
	@docker run --network=host port-service