package handler

import (
	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"github.com/YotoHana/itk-academy-test-case/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WalletHandler struct {
	walletService service.WalletService
}

func (h *WalletHandler) GetBalance(c *fiber.Ctx) error {
	walletUUIDStr := c.Params("WALLET_UUID")

	walletUUID, err := uuid.Parse(walletUUIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid path parameter")
	}

	wallet, err := h.walletService.GetWallet(c.Context(), walletUUID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(wallet)
}

func (h *WalletHandler) OperationWallet(c *fiber.Ctx) error {
	var req models.WalletRequest

	c.BodyParser(&req)

	wallet, err := h.walletService.OperateWallet(c.Context(), req)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(wallet)
}

func NewWalletHandler(walletService service.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}