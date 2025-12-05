package handler

import (
	"errors"
	"log"
	"time"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"github.com/YotoHana/itk-academy-test-case/internal/service"
	internalErrors "github.com/YotoHana/itk-academy-test-case/internal/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WalletHandler struct {
	walletService service.WalletService
}

func (h *WalletHandler) GetBalance(c *fiber.Ctx) error {
	start := time.Now()

	walletUUIDStr := c.Params("WALLET_UUID")

	walletUUID, err := uuid.Parse(walletUUIDStr)
	if err != nil {
		return checkError(c, start, err)
	}

	wallet, err := h.walletService.GetWallet(c.Context(), walletUUID)
	if err != nil {
		return checkError(c, start, err)
	}

	log.Printf("%s %s %d OK | time: %v", c.Method(), c.OriginalURL(), fiber.StatusOK, time.Since(start))
	return c.JSON(wallet)
}

func (h *WalletHandler) OperationWallet(c *fiber.Ctx) error {
	var req models.WalletRequest
	start := time.Now()

	c.BodyParser(&req)

	wallet, err := h.walletService.OperateWallet(c.Context(), req)
	if err != nil {
		return checkError(c, start, err)
	}

	log.Printf("%s %s %d OK | time: %v", c.Method(), c.OriginalURL(), fiber.StatusOK, time.Since(start))
	return c.JSON(wallet)
}

func NewWalletHandler(walletService service.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func checkError(c *fiber.Ctx, start time.Time, err error) error {
	switch {
	case errors.Is(err, internalErrors.ErrWalletNotFound):
		log.Printf("%s %s %d NotFound | time: %v", c.Method(), c.OriginalURL(), fiber.StatusNotFound, time.Since(start))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code": 404,
			"message": "wallet not found",
		})
	case errors.Is(err, internalErrors.ErrInvalidAmount):
		log.Printf("%s %s %d BadRequest | time: %v", c.Method(), c.OriginalURL(), fiber.StatusBadRequest, time.Since(start))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": 400,
			"message": "invalid amount",
		})
	case errors.Is(err, internalErrors.ErrInvalidOperationType):
		log.Printf("%s %s %d BadRequest | time: %v", c.Method(), c.OriginalURL(), fiber.StatusBadRequest, time.Since(start))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": 400,
			"message": "invalid operation type",
		})
	case errors.Is(err, internalErrors.ErrUnsufficientBalance):
		log.Printf("%s %s %d BadRequest | time: %v", c.Method(), c.OriginalURL(), fiber.StatusBadRequest, time.Since(start))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": 400,
			"message": "unsufficient balance",
		})
	case errors.Is(err, internalErrors.ErrInvalidPathParameter):
		log.Printf("%s %s %d BadRequest | time: %v", c.Method(), c.OriginalURL(), fiber.StatusBadRequest, time.Since(start))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": 400,
			"message": "invalid path parameter",
		})
	default:
		log.Printf("%s %s %d InternalServerError | time: %v", c.Method(), c.OriginalURL(), fiber.StatusInternalServerError, time.Since(start))
		return fiber.ErrInternalServerError
	}
}