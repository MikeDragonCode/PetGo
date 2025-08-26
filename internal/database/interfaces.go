package database

import "petProjectMike/internal/models"

// Database интерфейс для работы с базой данных
type Database interface {
	// Account operations
	CreateAccount(account *models.Account) error
	GetAccount(id string) (*models.Account, error)
	GetAccountsByUserID(userID string) ([]*models.Account, error)
	UpdateAccount(account *models.Account) error
	DeleteAccount(id string) error

	// Transaction operations
	CreateTransaction(transaction *models.Transaction) error
	GetTransaction(id string) (*models.Transaction, error)
	GetTransactionsByAccount(accountID string) ([]*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id string) error

	// Bonus operations
	CreateBonus(bonus *models.Bonus) error
	GetBonus(id string) (*models.Bonus, error)
	GetBonusesByUserID(userID string) ([]*models.Bonus, error)
	UpdateBonus(bonus *models.Bonus) error
	DeleteBonus(id string) error

	// User operations
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}
