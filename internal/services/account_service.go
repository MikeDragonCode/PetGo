package services

import (
	"errors"
	"time"

	"petProjectMike/internal/database"
	"petProjectMike/internal/models"
)

type AccountService struct {
	db database.Database
}

func NewAccountService(db database.Database) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateAccount(userID, currency string) (*models.Account, error) {
	_, err := s.db.GetUser(userID)
	if err != nil {
		return nil, err
	}

	supportedCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"RUB": true,
	}

	if !supportedCurrencies[currency] {
		return nil, errors.New("unsupported currency")
	}

	account := models.NewAccount(userID, currency)
	if err := s.db.CreateAccount(account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	return s.db.GetAccount(id)
}

func (s *AccountService) GetAccountsByUser(userID string) ([]*models.Account, error) {
	return s.db.GetAccountsByUserID(userID)
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	_, err := s.db.GetAccount(account.ID)
	if err != nil {
		return err
	}
	account.UpdatedAt = time.Now()
	return s.db.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id string) error {
	account, err := s.db.GetAccount(id)
	if err != nil {
		return err
	}
	if account.Balance > 0 {
		return errors.New("cannot delete account with positive balance")
	}
	return s.db.DeleteAccount(id)
}

func (s *AccountService) GetAccountBalance(id string) (float64, error) {
	account, err := s.db.GetAccount(id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

func (s *AccountService) ValidateAccount(id string) error {
	_, err := s.db.GetAccount(id)
	return err
}

func (s *AccountService) GetAccountSummary(id string) (map[string]interface{}, error) {
	account, err := s.db.GetAccount(id)
	if err != nil {
		return nil, err
	}
	transactions, err := s.db.GetTransactionsByAccount(id)
	if err != nil {
		return nil, err
	}
	summary := map[string]interface{}{
		"account_id":         account.ID,
		"user_id":            account.UserID,
		"balance":            account.Balance,
		"currency":           account.Currency,
		"created_at":         account.CreatedAt,
		"updated_at":         account.UpdatedAt,
		"total_transactions": len(transactions),
	}
	return summary, nil
}

func (s *AccountService) GetDB() database.Database {
	return s.db
}
