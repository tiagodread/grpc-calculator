package service

import (
	"context"
	"errors"

	pb "github.com/tiagodread/grpc-tickets/pkg/services/factors"
)

type CalculatorService interface {
	Sum(ctx context.Context, a, b int32, source string) (int32, error)
	Subtract(ctx context.Context, a, b int32, source string) (int32, error)
	Multiply(ctx context.Context, a, b int32, source string) (int32, error)
	Divide(ctx context.Context, a, b int32, source string) (int32, error)
}

type FactorClient interface {
	GetFactor(ctx context.Context, source string) (*pb.FactorResponse, error)
}

type calculatorServiceImpl struct {
	factorClient FactorClient
}

func NewCalculatorService(factorClient FactorClient) CalculatorService {
	return &calculatorServiceImpl{
		factorClient: factorClient,
	}
}

func (c *calculatorServiceImpl) Sum(ctx context.Context, a, b int32, source string) (int32, error) {
	factor, err := c.factorClient.GetFactor(ctx, source)
	if err != nil {
		return 0, err
	}
	result := a + b + factor.Value
	return result, nil
}

func (c *calculatorServiceImpl) Divide(ctx context.Context, a int32, b int32, source string) (int32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	factor, err := c.factorClient.GetFactor(ctx, source)
	if err != nil {
		return 0, err
	}
	result := a/b + factor.Value
	return result, nil
}

func (c *calculatorServiceImpl) Multiply(ctx context.Context, a int32, b int32, source string) (int32, error) {
	factor, err := c.factorClient.GetFactor(ctx, source)
	if err != nil {
		return 0, err
	}
	result := a*b + factor.Value
	return result, nil
}

func (c *calculatorServiceImpl) Subtract(ctx context.Context, a int32, b int32, source string) (int32, error) {
	factor, err := c.factorClient.GetFactor(ctx, source)
	if err != nil {
		return 0, err
	}
	result := a - b + factor.Value
	return result, nil
}
