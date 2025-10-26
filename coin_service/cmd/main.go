package main

import (
	"coin_service/config"
	"coin_service/internal"
	"coin_service/repository"
	"coin_service/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := sqlx.Open("postgres", cfg.Postgres.ConnectionURL())
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)

	r := gin.Default()
	ctrl := internal.NewController(r, svc)

	if err := ctrl.RunServer(":" + cfg.HTTPPort); err != nil {
		log.Fatal("Server error:", err)
	}
}
