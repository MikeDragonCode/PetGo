package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID          string    `json:"id"`
	FromAccount string    `json:"from_account"`
	ToAccount   string    `json:"to_account"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Bonus struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(userID, currency string) *Account {
	now := time.Now()
	return &Account{
		ID:        uuid.New().String(),
		UserID:    userID,
		Balance:   0.0,
		Currency:  currency,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewTransaction(fromAccount, toAccount string, amount float64, transactionType, description string) *Transaction {
	now := time.Now()
	return &Transaction{
		ID:          uuid.New().String(),
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Type:        transactionType,
		Status:      "pending",
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func NewBonus(userID, bonusType string, amount float64, expiresAt time.Time) *Bonus {
	now := time.Now()
	return &Bonus{
		ID:        uuid.New().String(),
		UserID:    userID,
		Type:      bonusType,
		Amount:    amount,
		Status:    "active",
		ExpiresAt: expiresAt,
		CreatedAt: now,
	}
}
