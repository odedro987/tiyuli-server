gen-proto: gen-proto-root gen-proto-expense

gen-proto-root:
	protoc \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto \
		--go-grpc_opt=paths=source_relative \
		--proto_path=./proto \
		--proto_path=./proto/expense \
		./proto/*.proto

gen-proto-expense:
	protoc \
		--go_out=./proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=./proto \
		--go-grpc_opt=paths=source_relative \
		--proto_path=./proto \
		--proto_path=./proto/expense \
		./proto/expense/*.proto