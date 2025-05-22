package main

import (
	"log"
	"net"

	pb "github.com/tiagodread/grpc-tickets/pkg/services/calculator"

	"github.com/tiagodread/grpc-tickets/services/calculator/internal/handler"
	"github.com/tiagodread/grpc-tickets/services/calculator/internal/service"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	calculadoraService := service.NewCalculatorService()
	calculadoraHandler := handler.NewCalculadoraHandler(calculadoraService)

	pb.RegisterCalculatorServiceServer(grpcServer, calculadoraHandler)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Default().Println("Server started on port 50051")
}
