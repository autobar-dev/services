run:
	go run main.go

build:
	go build -o bin/currency main.go

protos:
	protoc \
		--proto_path=../protos \
		--go_opt=Mcurrency.proto=./generated_grpc \
		--go-grpc_opt=Mcurrency.proto=./generated_grpc \
		--go_out=./grpc \
		--go-grpc_out=./grpc \
		currency.proto
