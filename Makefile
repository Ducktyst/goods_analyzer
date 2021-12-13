#PROTO_OUT_DIR := api/proto
PROTO_API_DIR := api

gen: proto gateway
	echo "Generate"

proto:
	protoc -I . \
	   --go_out . \
	   --go_opt paths=source_relative \
	   --go-grpc_out . \
	   --go-grpc_opt paths=source_relative \
	   ./$(PROTO_API_DIR)/price_analyze.proto

gateway:
	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./$(PROTO_API_DIR)/price_analyze.proto

migrate-run:
	@goose --dir="deployments/migrations" postgres "postgres://postgres:password@localhost:5432/catalog?sslmode=disable" up
