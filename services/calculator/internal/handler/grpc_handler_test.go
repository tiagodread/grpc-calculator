package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/tiagodread/grpc-tickets/pkg/services/calculator"
	"github.com/tiagodread/grpc-tickets/services/calculator/internal/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestSumHandler_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Sum(gomock.Any(), int32(1), int32(2)).
		Return(int32(3), nil)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Sum(ctx, &req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, int32(3), resp.Result)

}

func TestSumHandler_Error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Sum(gomock.Any(), int32(1), int32(2)).
		Return(int32(0), assert.AnError)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Sum(ctx, &req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestSubtractHandler_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Subtract(gomock.Any(), int32(1), int32(2)).
		Return(int32(-1), nil)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Subtract(ctx, &req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, int32(-1), resp.Result)
}
func TestSubtractHandler_Error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Subtract(gomock.Any(), int32(1), int32(2)).
		Return(int32(0), assert.AnError)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Subtract(ctx, &req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestMultiplyHandler_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Multiply(gomock.Any(), int32(1), int32(2)).
		Return(int32(2), nil)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Multiply(ctx, &req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, int32(2), resp.Result)
}
func TestMultiplyHandler_Error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Multiply(gomock.Any(), int32(1), int32(2)).
		Return(int32(0), assert.AnError)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Multiply(ctx, &req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestDivideHandler_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Divide(gomock.Any(), int32(1), int32(2)).
		Return(int32(0), nil)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Divide(ctx, &req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, int32(0), resp.Result)
}
func TestDivideHandler_Error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Divide(gomock.Any(), int32(1), int32(2)).
		Return(int32(0), assert.AnError)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 2,
	}

	resp, err := h.Divide(ctx, &req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestDivideByZeroHandler_Error(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Divide(gomock.Any(), int32(1), int32(0)).
		Return(int32(0), assert.AnError)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 0,
	}

	resp, err := h.Divide(ctx, &req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestDivideByZeroHandler_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCalculatorService := mocks.NewMockCalculatorService(ctrl)
	mockCalculatorService.EXPECT().
		Divide(gomock.Any(), int32(1), int32(0)).
		Return(int32(0), nil)

	h := NewCalculadoraHandler(mockCalculatorService)

	// act
	req := pb.Request{
		A: 1,
		B: 0,
	}

	resp, err := h.Divide(ctx, &req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, int32(0), resp.Result)
}
