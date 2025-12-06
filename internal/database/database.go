package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Database *gorm.DB
}

func NewDatabase(cfg *Config) (*Database, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		 cfg.User,
		 cfg.Password,
		 cfg.Host,
		 cfg.Port,
		 cfg.DBName,
		 cfg.SSLMode,
	)

	gorm, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{
		Database: gorm,
	}, nil
}