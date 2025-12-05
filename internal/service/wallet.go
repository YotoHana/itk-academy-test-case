package service

import (
	"context"
	"fmt"
	"log"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"github.com/YotoHana/itk-academy-test-case/internal/repository"
	"github.com/google/uuid"
)

type WalletService interface {
	GetWallet(ctx context.Context, uuid uuid.UUID) (*models.Wallets, error)
	OperateWallet(ctx context.Context, wallet models.WalletRequest) (*models.Wallets, error)
}

type walletService struct {
	walletRepo repository.WalletRepository
}

func (s *walletService) GetWallet(ctx context.Context, uuid uuid.UUID) (*models.Wallets, error) {
	wallet := &models.Wallets{
		WalletUUID: uuid,
	}

	if err := s.walletRepo.GetByID(ctx, wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *walletService) OperateWallet(ctx context.Context, req models.WalletRequest) (*models.Wallets, error) {
	log.Printf("req: %v", req)
	switch req.OperationType {
	case "DEPOSIT":
		wallet := &models.Wallets{
			WalletUUID: req.WalletUUID,
		}

		if err := s.walletRepo.GetByID(ctx, wallet); err != nil {
			return nil, err
		}

		wallet.Balance += req.Amount

		log.Printf("BALANCE: %d", wallet.Balance)

		if err := s.walletRepo.Update(ctx, wallet); err != nil {
			return nil, err
		}

		return wallet, nil
	
	case "WITHDRAW":
		wallet := &models.Wallets{
			WalletUUID: req.WalletUUID,
		}

		if err := s.walletRepo.GetByID(ctx, wallet); err != nil {
			return nil, err
		}

		if wallet.Balance - req.Amount >= 0 {
			wallet.Balance -= req.Amount

			if err := s.walletRepo.Update(ctx, wallet); err != nil {
				return nil, err
			}

			return wallet, nil
		}

		return nil, fmt.Errorf("Not enough balance for withdraw! Balance: %d, amount: %d", wallet.Balance, req.Amount)

	default:
		return nil, fmt.Errorf("Invalid operation type: %s", req.OperationType)
	}
}

func NewWalletService(walletRepo repository.WalletRepository) WalletService {
	return &walletService{walletRepo: walletRepo}
}
