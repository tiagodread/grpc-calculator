proto:
	protoc --go_out=pkg --go-grpc_out=pkg pkg/services/calculator/calculator.proto

mocks:
	mockgen \
  -source=services/calculator/internal/service/calculator_service.go \
  -destination=services/calculator/internal/service/mocks/mock_calculator_service.go \
  -package=mocks

grpc-server:
	@echo "Starting gRPC server..."
	@go run services/calculator/cmd/grpc/main.go
	@echo "gRPC server started."
	@echo "Listening on port 50051..."

test:
	@echo "Running tests..."
	@go test ./... -v --coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Tests completed."

build:
	@echo "Building the application..."
	@go build -o calculator ./services/calculator/cmd/grpc/main.go
	@echo "Build completed."
	@echo "Executable: ./calculator"