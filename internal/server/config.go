package server

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}, nil
}