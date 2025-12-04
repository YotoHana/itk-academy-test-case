package service

import (
	"context"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"github.com/YotoHana/itk-academy-test-case/internal/repository"
	"github.com/google/uuid"
)

type WalletService interface {
	GetWallet(ctx context.Context, uuid uuid.UUID) (*models.Wallets, error)
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

func NewWalletService(walletRepo repository.WalletRepository) WalletService {
	return &walletService{walletRepo: walletRepo}
}
