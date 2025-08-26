package services

import (
	"errors"
	"time"

	"petProjectMike/internal/database"
	"petProjectMike/internal/models"
)

type BonusService struct {
	db database.Database
}

func NewBonusService(db database.Database) *BonusService {
	return &BonusService{db: db}
}

func (s *BonusService) CreateWelcomeBonus(userID string, amount float64) (*models.Bonus, error) {
	_, err := s.db.GetUser(userID)
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().AddDate(0, 0, 30)
	bonus := models.NewBonus(userID, "welcome", amount, expiresAt)
	if err := s.db.CreateBonus(bonus); err != nil {
		return nil, err
	}
	return bonus, nil
}

func (s *BonusService) CreateTransactionBonus(userID string, amount float64, transactionType string) (*models.Bonus, error) {
	_, err := s.db.GetUser(userID)
	if err != nil {
		return nil, err
	}

	var bonusType string
	var bonusAmount float64
	switch transactionType {
	case "transfer":
		bonusType = "transaction"
		bonusAmount = amount * 0.01
	case "deposit":
		bonusType = "transaction"
		bonusAmount = amount * 0.005
	default:
		return nil, errors.New("unsupported transaction type for bonus")
	}

	expiresAt := time.Now().AddDate(0, 0, 90)
	bonus := models.NewBonus(userID, bonusType, bonusAmount, expiresAt)
	if err := s.db.CreateBonus(bonus); err != nil {
		return nil, err
	}
	return bonus, nil
}

func (s *BonusService) UseBonus(bonusID, accountID string) error {
	bonus, err := s.db.GetBonus(bonusID)
	if err != nil {
		return err
	}
	if bonus.Status != "active" {
		return errors.New("bonus is not active")
	}
	if time.Now().After(bonus.ExpiresAt) {
		bonus.Status = "expired"
		_ = s.db.UpdateBonus(bonus)
		return errors.New("bonus has expired")
	}

	account, err := s.db.GetAccount(accountID)
	if err != nil {
		return err
	}
	if account.UserID != bonus.UserID {
		return errors.New("bonus can only be used on user's own account")
	}

	account.Balance += bonus.Amount
	account.UpdatedAt = time.Now()
	if err := s.db.UpdateAccount(account); err != nil {
		return err
	}

	bonus.Status = "used"
	if err := s.db.UpdateBonus(bonus); err != nil {
		return err
	}
	return nil
}

func (s *BonusService) GetActiveBonuses(userID string) ([]*models.Bonus, error) {
	bonuses, err := s.db.GetBonusesByUserID(userID)
	if err != nil {
		return nil, err
	}
	var activeBonuses []*models.Bonus
	now := time.Now()
	for _, bonus := range bonuses {
		if bonus.Status == "active" && now.Before(bonus.ExpiresAt) {
			activeBonuses = append(activeBonuses, bonus)
		}
	}
	return activeBonuses, nil
}

func (s *BonusService) GetBonus(id string) (*models.Bonus, error) {
	return s.db.GetBonus(id)
}

func (s *BonusService) ExpireExpiredBonuses() error {
	return nil
}
