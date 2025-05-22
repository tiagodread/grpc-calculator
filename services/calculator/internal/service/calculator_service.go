package service

import (
	"context"
	"errors"
)

type CalculatorService interface {
	Sum(ctx context.Context, a, b int32) (int32, error)
	Subtract(ctx context.Context, a, b int32) (int32, error)
	Multiply(ctx context.Context, a, b int32) (int32, error)
	Divide(ctx context.Context, a, b int32) (int32, error)
}

type calculatorServiceImpl struct{}

func NewCalculatorService() CalculatorService {
	return &calculatorServiceImpl{}
}

func (c *calculatorServiceImpl) Sum(ctx context.Context, a, b int32) (int32, error) {
	result := a + b
	return result, nil
}

func (c *calculatorServiceImpl) Divide(ctx context.Context, a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	return result, nil
}

func (c *calculatorServiceImpl) Multiply(ctx context.Context, a int32, b int32) (int32, error) {
	result := a * b
	return result, nil
}

func (c *calculatorServiceImpl) Subtract(ctx context.Context, a int32, b int32) (int32, error) {
	result := a - b
	return result, nil
}
