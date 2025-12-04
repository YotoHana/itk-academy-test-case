package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YotoHana/itk-academy-test-case/internal/database"
	"github.com/YotoHana/itk-academy-test-case/internal/repository"
	"github.com/YotoHana/itk-academy-test-case/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	repo := repository.NewWalletRepository(db.Database)

	server := server.NewServer()

	go func ()  {
		<- ctx.Done()
		if err := server.Stop(context.Background()); err != nil {
			log.Printf("failed shutdown server: %v", err)
		}
	}()

	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}