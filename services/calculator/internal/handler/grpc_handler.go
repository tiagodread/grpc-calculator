package handler

import (
	"context"

	pb "github.com/tiagodread/grpc-tickets/pkg/services/calculator"
	"github.com/tiagodread/grpc-tickets/services/calculator/internal/service"
)

type CalculadoraHandler struct {
	pb.UnimplementedCalculatorServiceServer
	calculadoraService service.CalculatorService
}

func NewCalculadoraHandler(s service.CalculatorService) *CalculadoraHandler {
	return &CalculadoraHandler{
		calculadoraService: s,
	}
}

func (h *CalculadoraHandler) Sum(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result, err := h.calculadoraService.Sum(ctx, in.GetA(), in.GetB())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Result: result}, nil
}

func (h *CalculadoraHandler) Subtract(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result, err := h.calculadoraService.Subtract(ctx, in.GetA(), in.GetB())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Result: result}, nil
}

func (h *CalculadoraHandler) Divide(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result, err := h.calculadoraService.Divide(ctx, in.GetA(), in.GetB())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Result: result}, nil
}

func (h *CalculadoraHandler) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result, err := h.calculadoraService.Multiply(ctx, in.GetA(), in.GetB())
	if err != nil {
		return nil, err
	}
	return &pb.Response{Result: result}, nil
}
