package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	internalErrors "github.com/YotoHana/itk-academy-test-case/internal/errors"
	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock для WalletService
type MockWalletService struct {
	mock.Mock
}

func (m *MockWalletService) GetWallet(ctx context.Context, walletUUID uuid.UUID) (*models.Wallets, error) {
	args := m.Called(ctx, walletUUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Wallets), args.Error(1)
}

func (m *MockWalletService) OperateWallet(ctx context.Context, req models.WalletRequest) (*models.Wallets, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Wallets), args.Error(1)
}

// Тесты для GetBalance

func TestGetBalance_Success(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	expectedWallet := &models.Wallets{
		WalletUUID: walletUUID,
		Balance:    5000,
	}

	mockService.On("GetWallet", mock.Anything, walletUUID).Return(expectedWallet, nil)

	app.Get("/api/v1/wallets/:WALLET_UUID", handler.GetBalance)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/wallets/%s", walletUUID.String()), nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var wallet models.Wallets
	json.Unmarshal(body, &wallet)

	assert.Equal(t, walletUUID, wallet.WalletUUID)
	assert.Equal(t, 5000, wallet.Balance)

	mockService.AssertExpectations(t)
}

func TestGetBalance_InvalidUUID(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	app.Get("/api/v1/wallets/:WALLET_UUID", handler.GetBalance)

	req := httptest.NewRequest("GET", "/api/v1/wallets/invalid-uuid", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]any
	json.Unmarshal(body, &response)

	assert.Equal(t, float64(400), response["code"])
	assert.Equal(t, "invalid path parameter", response["message"])

	mockService.AssertNotCalled(t, "GetWallet")
}

func TestGetBalance_WalletNotFound(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()

	mockService.On("GetWallet", mock.Anything, walletUUID).Return(nil, internalErrors.ErrWalletNotFound)

	app.Get("/api/v1/wallets/:WALLET_UUID", handler.GetBalance)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/wallets/%s", walletUUID.String()), nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	assert.Equal(t, float64(404), response["code"])
	assert.Equal(t, "wallet not found", response["message"])

	mockService.AssertExpectations(t)
}

// Тесты для OperationWallet

func TestOperationWallet_Deposit_Success(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "DEPOSIT",
		Amount:        1000,
	}

	expectedWallet := &models.Wallets{
		WalletUUID: walletUUID,
		Balance:    6000,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(expectedWallet, nil)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var wallet models.Wallets
	json.Unmarshal(body, &wallet)

	assert.Equal(t, walletUUID, wallet.WalletUUID)
	assert.Equal(t, 6000, wallet.Balance)

	mockService.AssertExpectations(t)
}

func TestOperationWallet_Withdraw_Success(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "WITHDRAW",
		Amount:        500,
	}

	expectedWallet := &models.Wallets{
		WalletUUID: walletUUID,
		Balance:    4500,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(expectedWallet, nil)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestOperationWallet_InvalidAmount(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "DEPOSIT",
		Amount:        0,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(nil, internalErrors.ErrInvalidAmount)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	assert.Equal(t, float64(400), response["code"])
	assert.Equal(t, "invalid amount", response["message"])

	mockService.AssertExpectations(t)
}

func TestOperationWallet_InsufficientBalance(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "WITHDRAW",
		Amount:        10000,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(nil, internalErrors.ErrUnsufficientBalance)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	assert.Equal(t, float64(400), response["code"])
	assert.Equal(t, "unsufficient balance", response["message"])

	mockService.AssertExpectations(t)
}

func TestOperationWallet_InvalidOperationType(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "INVALID",
		Amount:        1000,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(nil, internalErrors.ErrInvalidOperationType)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	assert.Equal(t, float64(400), response["code"])
	assert.Equal(t, "invalid operation type", response["message"])

	mockService.AssertExpectations(t)
}

func TestOperationWallet_WalletNotFound(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "DEPOSIT",
		Amount:        1000,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(nil, internalErrors.ErrWalletNotFound)

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestOperationWallet_InternalServerError(t *testing.T) {
	app := fiber.New()
	mockService := new(MockWalletService)
	handler := NewWalletHandler(mockService)

	walletUUID := uuid.New()
	requestBody := models.WalletRequest{
		WalletUUID:    walletUUID,
		OperationType: "DEPOSIT",
		Amount:        1000,
	}

	mockService.On("OperateWallet", mock.Anything, requestBody).Return(nil, errors.New("database error"))

	app.Post("/api/v1/wallet", handler.OperationWallet)

	bodyBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	mockService.AssertExpectations(t)
}