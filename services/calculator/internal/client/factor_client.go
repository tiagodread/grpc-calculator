package client

import (
	"context"
	"log"

	pb "github.com/tiagodread/grpc-tickets/pkg/services/factors"
	"google.golang.org/grpc"
)

type FactorClient struct {
	conn   *grpc.ClientConn
	client pb.FactorsServiceClient
}

func NewFactorClient(address string) (*FactorClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &FactorClient{
		conn:   conn,
		client: pb.NewFactorsServiceClient(conn),
	}, nil
}

func (fc *FactorClient) GetFactor(ctx context.Context, source string) (*pb.FactorResponse, error) {
	req := &pb.FactorRequest{
		Source: source,
	}

	resp, err := fc.client.GetFactor(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (fc *FactorClient) Close() {

	if err := fc.conn.Close(); err != nil {
		log.Printf("erro ao fechar conexão gRPC: %v", err)
	}
}
