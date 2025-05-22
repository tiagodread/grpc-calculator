package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// arrange
	calculatorService := NewCalculatorService()
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		expected int32
	}{
		{1, 2, 3},
		{2, 3, 5},
		{5, 5, 10},
		{10, 20, 30},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Sum(ctx, c.a, c.b)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}

}

func TestSubtract(t *testing.T) {
	// arrange
	calculatorService := NewCalculatorService()
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		expected int32
	}{
		{1, 2, -1},
		{2, 3, -1},
		{5, 5, 0},
		{10, 20, -10},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Subtract(ctx, c.a, c.b)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}

}
func TestMultiply(t *testing.T) {
	// arrange
	calculatorService := NewCalculatorService()
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		expected int32
	}{
		{1, 2, 2},
		{2, 3, 6},
		{5, 5, 25},
		{10, 20, 200},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Multiply(ctx, c.a, c.b)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}

}
func TestDivide(t *testing.T) {
	// arrange
	calculatorService := NewCalculatorService()
	ctx := context.Background()
	cases := []struct {
		a, b     int32
		expected int32
	}{
		{1, 2, 0},
		{2, 3, 0},
		{5, 5, 1},
		{10, 20, 0},
	}

	// act
	for _, c := range cases {
		result, err := calculatorService.Divide(ctx, c.a, c.b)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, result, c.expected)
	}

}
func TestDivideByZero(t *testing.T) {
	// arrange
	calculatorService := NewCalculatorService()
	ctx := context.Background()

	// act
	result, err := calculatorService.Divide(ctx, 10, 0)

	// assert
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "division by zero")
	assert.Equal(t, result, int32(0))
}
