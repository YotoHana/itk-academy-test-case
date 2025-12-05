package repository

import (
	"context"

	"github.com/YotoHana/itk-academy-test-case/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Update(ctx context.Context, wallet *models.Wallets) error
	GetByID(ctx context.Context, wallet *models.Wallets) error
}

type walletRepo struct {
	db *gorm.DB
}

func (r *walletRepo) Update(ctx context.Context, wallet *models.Wallets) error {
	return r.db.WithContext(ctx).Model(&models.Wallets{}).Where("wallet_uuid = ?", wallet.WalletUUID).Updates(&wallet).Error
}

func (r *walletRepo) GetByID(ctx context.Context, wallet *models.Wallets) error {
	return r.db.WithContext(ctx).Model(&models.Wallets{}).First(&wallet, wallet.WalletUUID).Error
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepo{db: db}
}