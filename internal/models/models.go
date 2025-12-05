package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallets struct {
	WalletUUID uuid.UUID `gorm:"primaryKey;column:wallet_uuid;type:uuid"`
	Balance int `gorm:"column:balance;type:integer;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT TIMESTAMP"`
	UpdatedAt time.Time `gotm:"column:updated_at;type:timestamp"`
}

type WalletRequest struct {
	WalletUUID uuid.UUID `json:"wallet_id"`
	OperationType string `json:"operation_type"`
	Amount int `json:"amount"`
}