package server

import (
	"context"
	"fmt"

	"github.com/YotoHana/itk-academy-test-case/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app     *fiber.App
	handler *handler.WalletHandler
	cfg     *Config
}

func (s *Server) Start() error {
	listenAddr := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	return s.app.Listen(listenAddr)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) registerRoutes() {
	api := s.app.Group("/api/v1/")
	api.Get("wallets/:WALLET_UUID", s.handler.GetBalance)
	api.Post("wallet", s.handler.OperationWallet)
}

func NewServer(handler *handler.WalletHandler, cfg *Config) *Server {
	app := fiber.New()

	s := &Server{
		app:     app,
		handler: handler,
		cfg:     cfg,
	}

	s.registerRoutes()

	return s
}
