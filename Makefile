#PROTO_OUT_DIR := price_analyzer_api/proto
PROTO_API_DIR := api

run:
	@echo "Compiling"
	@go run -mod=mod  cmd/price_analyzer/main.go

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
		./$(PROTO_API_DIR)/price_analyze.proto

migrate-run:
	@goose --dir="deployments/migrations" postgres "postgres://postgres:password@localhost:5432/price_analyze?sslmode=disable" up
