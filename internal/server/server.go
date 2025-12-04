package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func (s *Server) Start() error {
	return s.app.Listen(":8080")
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func NewServer() *Server {
	app := fiber.New()

	return &Server{
		app: app,
	}
}