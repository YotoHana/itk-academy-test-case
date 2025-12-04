package server

import (
	"context"

	"github.com/YotoHana/itk-academy-test-case/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	handler *handler.WalletHandler
}

func (s *Server) Start() error {
	return s.app.Listen(":8080")
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) registerRoutes() {
	api := s.app.Group("/api/v1/")
	api.Get("wallets/:WALLET_UUID", s.handler.GetBalance)
}

func NewServer(handler *handler.WalletHandler) *Server {
	app := fiber.New()

	s := &Server{
		app: app,
		handler: handler,
	}

	s.registerRoutes()

	return s
}