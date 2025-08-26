package main

import (
	"log"
	"os"

	"petProjectMike/internal/api"
	"petProjectMike/internal/config"
	"petProjectMike/internal/database"
	"petProjectMike/internal/services"
)

func main() {

	cfg := config.Load()

	db := database.NewInMemoryDB()

	transactionService := services.NewTransactionService(db)
	bonusService := services.NewBonusService(db)
	accountService := services.NewAccountService(db)

	server := api.NewServer(cfg, transactionService, bonusService, accountService)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
		os.Exit(1)
	}
}
