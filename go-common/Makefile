gen-proto: gen-proto-error

gen-proto-error:
	protoc \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto \
		--go-grpc_opt=paths=source_relative \
		--proto_path=./proto \
		--proto_path=./proto/error \
		./proto/error/*.proto