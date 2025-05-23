package main

import (
	"log"
	"net"

	pb "github.com/tiagodread/grpc-tickets/pkg/services/calculator"

	"github.com/tiagodread/grpc-tickets/services/calculator/internal/client"
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

	factorsClient, err := client.NewFactorClient("localhost:50052")
	if err != nil {
		log.Fatalf("failed to create factor client: %v", err)
	}
	defer factorsClient.Close()

	calculadoraService := service.NewCalculatorService(factorsClient)
	calculadoraHandler := handler.NewCalculadoraHandler(calculadoraService)

	pb.RegisterCalculatorServiceServer(grpcServer, calculadoraHandler)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Default().Println("Server started on port 50051")
}
