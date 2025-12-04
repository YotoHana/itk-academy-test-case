package repository

import (
	"context"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Update(ctx context.Context, balance int) (*models.Wallets, error)
}

type walletRepo struct {
	db *gorm.DB
}

func (r *walletRepo) Update(ctx context.Context, wallets models.Wallets) error {
	return r.db.Model(&models.Wallets{}).Select("wallet_uuid").Updates(&wallets).Error
}