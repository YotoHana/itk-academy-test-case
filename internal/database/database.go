package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	database *gorm.DB
}

func NewDatabase() (*Database, error) {
	gorm, err := gorm.Open(postgres.Open("postgres://dev:dev@localhost:5432/dev?sslmode=disable"))
	if err != nil {
		return nil, err
	}

	return &Database{
		database: gorm,
	}, nil
}