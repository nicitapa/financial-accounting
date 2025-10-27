package service

import (
	"coin_service/models"
	"coin_service/repository"
)

type TransactionService struct {
	Repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{Repo: repo}
}

func (s *TransactionService) CreateTransactions(tx models.Transaction) error {
	return s.Repo.CreateTransactions(tx)
}

func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.Repo.GetAllTransactions()

}
