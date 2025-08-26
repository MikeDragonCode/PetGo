package services

import (
	"testing"
	"time"

	"petProjectMike/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_CreateAccount(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		currency      string
		setupMocks    func(*MockDatabase)
		expectedError bool
	}{
		{
			name:     "successful account creation USD",
			userID:   "user-1",
			currency: "USD",
			setupMocks: func(mockDB *MockDatabase) {
				user := &models.User{
					ID:    "user-1",
					Email: "test@example.com",
					Name:  "Test User",
				}
				mockDB.On("GetUser", "user-1").Return(user, nil)
				mockDB.On("CreateAccount", mock.AnythingOfType("*models.Account")).Return(nil)
			},
			expectedError: false,
		},
		{
			name:     "successful account creation EUR",
			userID:   "user-1",
			currency: "EUR",
			setupMocks: func(mockDB *MockDatabase) {
				user := &models.User{
					ID:    "user-1",
					Email: "test@example.com",
					Name:  "Test User",
				}
				mockDB.On("GetUser", "user-1").Return(user, nil)
				mockDB.On("CreateAccount", mock.AnythingOfType("*models.Account")).Return(nil)
			},
			expectedError: false,
		},
		{
			name:     "unsupported currency",
			userID:   "user-1",
			currency: "GBP",
			setupMocks: func(mockDB *MockDatabase) {
				user := &models.User{
					ID:    "user-1",
					Email: "test@example.com",
					Name:  "Test User",
				}
				mockDB.On("GetUser", "user-1").Return(user, nil)
			},
			expectedError: true,
		},
		{
			name:     "user not found",
			userID:   "user-999",
			currency: "USD",
			setupMocks: func(mockDB *MockDatabase) {
				mockDB.On("GetUser", "user-999").Return(nil, assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDatabase{}
			tt.setupMocks(mockDB)

			service := NewAccountService(mockDB)
			account, err := service.CreateAccount(tt.userID, tt.currency)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, account)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, account)
				assert.Equal(t, tt.userID, account.UserID)
				assert.Equal(t, tt.currency, account.Currency)
				assert.Equal(t, 0.0, account.Balance)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestAccountService_GetAccount(t *testing.T) {
	mockDB := &MockDatabase{}
	expectedAccount := &models.Account{
		ID:       "account-1",
		UserID:   "user-1",
		Balance:  1000.0,
		Currency: "USD",
	}

	mockDB.On("GetAccount", "account-1").Return(expectedAccount, nil)

	service := NewAccountService(mockDB)
	account, err := service.GetAccount("account-1")

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)
	mockDB.AssertExpectations(t)
}

func TestAccountService_GetAccountsByUser(t *testing.T) {
	mockDB := &MockDatabase{}
	expectedAccounts := []*models.Account{
		{
			ID:       "account-1",
			UserID:   "user-1",
			Balance:  1000.0,
			Currency: "USD",
		},
		{
			ID:       "account-2",
			UserID:   "user-1",
			Balance:  500.0,
			Currency: "EUR",
		},
	}

	mockDB.On("GetAccountsByUserID", "user-1").Return(expectedAccounts, nil)

	service := NewAccountService(mockDB)
	accounts, err := service.GetAccountsByUser("user-1")

	assert.NoError(t, err)
	assert.Equal(t, expectedAccounts, accounts)
	mockDB.AssertExpectations(t)
}

func TestAccountService_UpdateAccount(t *testing.T) {
	mockDB := &MockDatabase{}
	account := &models.Account{
		ID:       "account-1",
		UserID:   "user-1",
		Balance:  1500.0,
		Currency: "USD",
	}

	mockDB.On("GetAccount", "account-1").Return(account, nil)
	mockDB.On("UpdateAccount", mock.AnythingOfType("*models.Account")).Return(nil)

	service := NewAccountService(mockDB)
	err := service.UpdateAccount(account)

	assert.NoError(t, err)
	// Проверяем, что время обновления было изменено
	assert.True(t, account.UpdatedAt.After(time.Now().Add(-time.Second)))
	mockDB.AssertExpectations(t)
}

func TestAccountService_DeleteAccount(t *testing.T) {
	tests := []struct {
		name          string
		accountID     string
		setupMocks    func(*MockDatabase)
		expectedError bool
	}{
		{
			name:      "successful deletion of empty account",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				account := &models.Account{
					ID:       "account-1",
					UserID:   "user-1",
					Balance:  0.0,
					Currency: "USD",
				}
				mockDB.On("GetAccount", "account-1").Return(account, nil)
				mockDB.On("DeleteAccount", "account-1").Return(nil)
			},
			expectedError: false,
		},
		{
			name:      "cannot delete account with positive balance",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				account := &models.Account{
					ID:       "account-1",
					UserID:   "user-1",
					Balance:  100.0,
					Currency: "USD",
				}
				mockDB.On("GetAccount", "account-1").Return(account, nil)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDatabase{}
			tt.setupMocks(mockDB)

			service := NewAccountService(mockDB)
			err := service.DeleteAccount(tt.accountID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestAccountService_GetAccountBalance(t *testing.T) {
	mockDB := &MockDatabase{}
	expectedAccount := &models.Account{
		ID:       "account-1",
		UserID:   "user-1",
		Balance:  1000.0,
		Currency: "USD",
	}

	mockDB.On("GetAccount", "account-1").Return(expectedAccount, nil)

	service := NewAccountService(mockDB)
	balance, err := service.GetAccountBalance("account-1")

	assert.NoError(t, err)
	assert.Equal(t, 1000.0, balance)
	mockDB.AssertExpectations(t)
}

func TestAccountService_ValidateAccount(t *testing.T) {
	mockDB := &MockDatabase{}
	expectedAccount := &models.Account{
		ID:       "account-1",
		UserID:   "user-1",
		Balance:  1000.0,
		Currency: "USD",
	}

	mockDB.On("GetAccount", "account-1").Return(expectedAccount, nil)

	service := NewAccountService(mockDB)
	err := service.ValidateAccount("account-1")

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestAccountService_GetAccountSummary(t *testing.T) {
	mockDB := &MockDatabase{}
	account := &models.Account{
		ID:        "account-1",
		UserID:    "user-1",
		Balance:   1000.0,
		Currency:  "USD",
		CreatedAt: time.Now().AddDate(0, 0, -1),
		UpdatedAt: time.Now(),
	}
	transactions := []*models.Transaction{
		{
			ID:          "txn-1",
			FromAccount: "account-1",
			ToAccount:   "account-2",
			Amount:      100.0,
			Type:        "transfer",
			Status:      "completed",
		},
		{
			ID:        "txn-2",
			ToAccount: "account-1",
			Amount:    50.0,
			Type:      "deposit",
			Status:    "completed",
		},
	}

	mockDB.On("GetAccount", "account-1").Return(account, nil)
	mockDB.On("GetTransactionsByAccount", "account-1").Return(transactions, nil)

	service := NewAccountService(mockDB)
	summary, err := service.GetAccountSummary("account-1")

	assert.NoError(t, err)
	assert.NotNil(t, summary)
	assert.Equal(t, "account-1", summary["account_id"])
	assert.Equal(t, "user-1", summary["user_id"])
	assert.Equal(t, 1000.0, summary["balance"])
	assert.Equal(t, "USD", summary["currency"])
	assert.Equal(t, 2, summary["total_transactions"])
	mockDB.AssertExpectations(t)
}
