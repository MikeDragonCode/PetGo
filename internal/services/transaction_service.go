package services

import (
	"errors"
	"time"

	"petProjectMike/internal/database"
	"petProjectMike/internal/models"
)

type TransactionService struct {
	db database.Database
}

func NewTransactionService(db database.Database) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) CreateTransfer(fromAccountID, toAccountID string, amount float64, description string) (*models.Transaction, error) {
	fromAccount, err := s.db.GetAccount(fromAccountID)
	if err != nil {
		return nil, err
	}
	toAccount, err := s.db.GetAccount(toAccountID)
	if err != nil {
		return nil, err
	}
	if fromAccount.Balance < amount {
		return nil, errors.New("insufficient funds")
	}
	if fromAccount.Currency != toAccount.Currency {
		return nil, errors.New("currency mismatch")
	}

	transaction := models.NewTransaction(fromAccountID, toAccountID, amount, "transfer", description)
	if err := s.db.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	fromAccount.Balance -= amount
	fromAccount.UpdatedAt = time.Now()
	if err := s.db.UpdateAccount(fromAccount); err != nil {
		return nil, err
	}

	toAccount.Balance += amount
	toAccount.UpdatedAt = time.Now()
	if err := s.db.UpdateAccount(toAccount); err != nil {
		return nil, err
	}

	transaction.Status = "completed"
	transaction.UpdatedAt = time.Now()
	if err := s.db.UpdateTransaction(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) CreateDeposit(accountID string, amount float64, description string) (*models.Transaction, error) {
	account, err := s.db.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	transaction := models.NewTransaction("", accountID, amount, "deposit", description)
	if err := s.db.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	account.Balance += amount
	account.UpdatedAt = time.Now()
	if err := s.db.UpdateAccount(account); err != nil {
		return nil, err
	}

	transaction.Status = "completed"
	transaction.UpdatedAt = time.Now()
	if err := s.db.UpdateTransaction(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) CreateWithdrawal(accountID string, amount float64, description string) (*models.Transaction, error) {
	account, err := s.db.GetAccount(accountID)
	if err != nil {
		return nil, err
	}
	if account.Balance < amount {
		return nil, errors.New("insufficient funds")
	}

	transaction := models.NewTransaction(accountID, "", amount, "withdrawal", description)
	if err := s.db.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	account.Balance -= amount
	account.UpdatedAt = time.Now()
	if err := s.db.UpdateAccount(account); err != nil {
		return nil, err
	}

	transaction.Status = "completed"
	transaction.UpdatedAt = time.Now()
	if err := s.db.UpdateTransaction(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactionHistory(accountID string) ([]*models.Transaction, error) {
	return s.db.GetTransactionsByAccount(accountID)
}

func (s *TransactionService) GetTransaction(id string) (*models.Transaction, error) {
	return s.db.GetTransaction(id)
}
