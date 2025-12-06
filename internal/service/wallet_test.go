package service

import (
	"context"
	"errors"
	"testing"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	internalErrors "github.com/YotoHana/itk-academy-test-case/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Update(ctx context.Context, wallet *models.Wallets) error {
	args := m.Called(ctx, wallet)
	return args.Error(0)
}

func (m *MockWalletRepository) GetByID(ctx context.Context, wallet *models.Wallets) error {
	args := m.Called(ctx, wallet)
	return args.Error(0)
}

func TestGetWallet_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	ctx := context.Background()
	uuid := uuid.New()

	mockRepo.On("GetByID", mock.Anything, mock.AnythingOfType("*models.Wallets")).Return(nil)

	wallet, err := service.GetWallet(ctx, uuid)

	assert.NoError(t, err)
	assert.NotNil(t, wallet)
	assert.Equal(t, uuid, wallet.WalletUUID)

	mockRepo.AssertExpectations(t)
}

func TestGetWallet_NotFound(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	ctx := context.Background()
	uuid := uuid.New()
	notFoundError := errors.New("not found")

	mockRepo.On("GetByID", mock.Anything, mock.AnythingOfType("*models.Wallets")).Return(notFoundError)

	wallet, err := service.GetWallet(ctx, uuid)

	assert.Error(t, err)
	assert.Nil(t, wallet)
	assert.Equal(t, internalErrors.ErrWalletNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestOperateWallet_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	ctx := context.Background()
	req := models.WalletRequest{
		WalletUUID: uuid.New(),
		OperationType: "DEPOSIT",
		Amount: 1000,
	}

	mockRepo.On("GetByID", mock.Anything, mock.AnythingOfType("*models.Wallets")).Return(nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.Wallets")).Return(nil)

	wallet, err := service.OperateWallet(ctx, req)

	assert.NotNil(t, wallet)
	assert.NoError(t, err)
	assert.Equal(t, req.WalletUUID, wallet.WalletUUID)

	mockRepo.AssertExpectations(t)
}

func TestOperateWallet_ZeroAmount(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	ctx := context.Background()
	req := models.WalletRequest{
		WalletUUID: uuid.New(),
		OperationType: "DEPOSIT",
		Amount: 0,
	}

	wallet, err := service.OperateWallet(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, wallet)
	assert.Equal(t, internalErrors.ErrInvalidAmount, err)

	mockRepo.AssertNotCalled(t, "GetByID")
	mockRepo.AssertNotCalled(t, "Update")
}

func TestOperateWallet_WrongOperationType(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	ctx := context.Background()
	req := models.WalletRequest{
		WalletUUID: uuid.New(),
		OperationType: "WRONG",
		Amount: 1000,
	}

	wallet, err := service.OperateWallet(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, wallet)
	assert.Equal(t, internalErrors.ErrInvalidOperationType, err)

	mockRepo.AssertNotCalled(t, "GetByID")
	mockRepo.AssertNotCalled(t, "Update")
}