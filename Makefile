.PHONY: docker down proto-gen grpc-test

docker:
	 docker-compose up --build -d

down:
	docker-compose down

proto-gen:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ ./proto/order.proto

grpc-test:
	grpcurl -plaintext localhost:50051 order.OrderService/ListOrders