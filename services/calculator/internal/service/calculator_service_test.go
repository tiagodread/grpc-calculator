package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagodread/grpc-tickets/services/calculator/internal/client"
)

func TestSum(t *testing.T) {
	// arrange
	factorClient, _ := client.NewFactorClient("localhost:50052")
	calculatorService := NewCalculatorService(factorClient)
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		source   string
		expected int32
	}{
		{1, 2, "BONUS", 5},
		{2, 3, "BONUS", 7},
		{5, 5, "BONUS", 12},
		{10, 20, "BONUS", 32},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Sum(ctx, c.a, c.b, c.source)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}

}

func TestSubtract(t *testing.T) {
	// arrange
	factorClient, _ := client.NewFactorClient("localhost:50052")
	calculatorService := NewCalculatorService(factorClient)
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		source   string
		expected int32
	}{
		{1, 2, "BONUS", 1},
		{2, 3, "BONUS", 1},
		{5, 5, "BONUS", 2},
		{10, 20, "BONUS", -8},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Subtract(ctx, c.a, c.b, c.source)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}
}

func TestMultiply(t *testing.T) {
	// arrange
	factorClient, _ := client.NewFactorClient("localhost:50052")
	calculatorService := NewCalculatorService(factorClient)
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		source   string
		expected int32
	}{
		{1, 2, "BONUS", 4},
		{2, 3, "BONUS", 8},
		{5, 5, "BONUS", 27},
		{10, 20, "BONUS", 202},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Multiply(ctx, c.a, c.b, c.source)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}
}

func TestDivide(t *testing.T) {
	// arrange
	factorClient, _ := client.NewFactorClient("localhost:50052")
	calculatorService := NewCalculatorService(factorClient)
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		source   string
		expected int32
	}{
		{1, 2, "BONUS", 2},
		{2, 3, "BONUS", 2},
		{5, 5, "BONUS", 3},
		{10, 20, "BONUS", 2},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Divide(ctx, c.a, c.b, c.source)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}
}

func TestDivideByZero(t *testing.T) {
	// arrange
	factorClient, _ := client.NewFactorClient("localhost:50052")
	calculatorService := NewCalculatorService(factorClient)
	ctx := context.Background()

	// act
	result, err := calculatorService.Divide(ctx, 10, 0, "BONUS")

	// assert
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "division by zero")
	assert.Equal(t, result, int32(0))
}
