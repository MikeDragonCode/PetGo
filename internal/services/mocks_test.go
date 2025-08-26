package services

import (
	"petProjectMike/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockDatabase реализует интерфейс Database для тестов сервисов
// Размещен в пакете services, чтобы быть доступным всем *_test.go в этом пакете
// и не экспортироваться наружу.
type MockDatabase struct {
	mock.Mock
}

// Account operations
func (m *MockDatabase) CreateAccount(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockDatabase) GetAccount(id string) (*models.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockDatabase) GetAccountsByUserID(userID string) ([]*models.Account, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Account), args.Error(1)
}

func (m *MockDatabase) UpdateAccount(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockDatabase) DeleteAccount(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Transaction operations
func (m *MockDatabase) CreateTransaction(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockDatabase) GetTransaction(id string) (*models.Transaction, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockDatabase) GetTransactionsByAccount(accountID string) ([]*models.Transaction, error) {
	args := m.Called(accountID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Transaction), args.Error(1)
}

func (m *MockDatabase) UpdateTransaction(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockDatabase) DeleteTransaction(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Bonus operations
func (m *MockDatabase) CreateBonus(bonus *models.Bonus) error {
	args := m.Called(bonus)
	return args.Error(0)
}

func (m *MockDatabase) GetBonus(id string) (*models.Bonus, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Bonus), args.Error(1)
}

func (m *MockDatabase) GetBonusesByUserID(userID string) ([]*models.Bonus, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Bonus), args.Error(1)
}

func (m *MockDatabase) UpdateBonus(bonus *models.Bonus) error {
	args := m.Called(bonus)
	return args.Error(0)
}

func (m *MockDatabase) DeleteBonus(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// User operations
func (m *MockDatabase) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDatabase) GetUser(id string) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockDatabase) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockDatabase) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDatabase) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
