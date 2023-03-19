run:
	go run main.go

build:
	go build -o bin/server

protos:
	protoc \
		--proto_path=../protos \
		--go_opt=Mcurrency.proto=./generated_grpc \
		--go-grpc_opt=Mcurrency.proto=./generated_grpc \
		--go_out=./grpc \
		--go-grpc_out=./grpc \
		currency.proto

swagger:
	swagger generate spec -o ../docs/currency/swagger.json --scan-models

migrations:
	migrate -database ${DB_CONNECTION_STRING} -path migrations up

new_migration:
	if [ -z "$(name)" ]; then echo "Provide name for migration as a parameter (make new_migration name=my_migration)"; else migrate create -ext sql -dir migrations -seq "$(name)"; fi
