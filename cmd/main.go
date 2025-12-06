package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YotoHana/itk-academy-test-case/internal/database"
	"github.com/YotoHana/itk-academy-test-case/internal/handler"
	"github.com/YotoHana/itk-academy-test-case/internal/repository"
	"github.com/YotoHana/itk-academy-test-case/internal/server"
	"github.com/YotoHana/itk-academy-test-case/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbCfg, err := database.NewConfig()
	if err != nil {
		fmt.Printf("failed to create db config: %v", err)
	}

	srvCfg, err := server.NewConfig()
	if err != nil {
		fmt.Printf("failed to create server config: %v", err)
	}

	db, err := database.NewDatabase(dbCfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	repo := repository.NewWalletRepository(db.Database)
	service := service.NewWalletService(repo)
	handlers := handler.NewWalletHandler(service)

	server := server.NewServer(handlers, srvCfg)

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