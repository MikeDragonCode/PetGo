package services

import (
	"testing"
	"time"

	"petProjectMike/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBonusService_CreateWelcomeBonus(t *testing.T) {
	mockDB := &MockDatabase{}
	user := &models.User{
		ID:    "user-1",
		Email: "test@example.com",
		Name:  "Test User",
	}

	mockDB.On("GetUser", "user-1").Return(user, nil)
	mockDB.On("CreateBonus", mock.AnythingOfType("*models.Bonus")).Return(nil)

	service := NewBonusService(mockDB)
	bonus, err := service.CreateWelcomeBonus("user-1", 50.0)

	assert.NoError(t, err)
	assert.NotNil(t, bonus)
	assert.Equal(t, "user-1", bonus.UserID)
	assert.Equal(t, "welcome", bonus.Type)
	assert.Equal(t, 50.0, bonus.Amount)
	assert.Equal(t, "active", bonus.Status)

	// Проверяем, что срок действия установлен на 30 дней вперед
	expectedExpiry := time.Now().AddDate(0, 0, 30)
	assert.WithinDuration(t, expectedExpiry, bonus.ExpiresAt, 2*time.Second)

	mockDB.AssertExpectations(t)
}

func TestBonusService_CreateTransactionBonus(t *testing.T) {
	tests := []struct {
		name            string
		userID          string
		amount          float64
		transactionType string
		expectedAmount  float64
		expectedError   bool
	}{
		{
			name:            "transfer bonus",
			userID:          "user-1",
			amount:          100.0,
			transactionType: "transfer",
			expectedAmount:  1.0, // 1% от 100
			expectedError:   false,
		},
		{
			name:            "deposit bonus",
			userID:          "user-1",
			amount:          200.0,
			transactionType: "deposit",
			expectedAmount:  1.0, // 0.5% от 200
			expectedError:   false,
		},
		{
			name:            "unsupported transaction type",
			userID:          "user-1",
			amount:          100.0,
			transactionType: "withdrawal",
			expectedError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDatabase{}
			user := &models.User{
				ID:    tt.userID,
				Email: "test@example.com",
				Name:  "Test User",
			}

			mockDB.On("GetUser", tt.userID).Return(user, nil)

			if !tt.expectedError {
				mockDB.On("CreateBonus", mock.AnythingOfType("*models.Bonus")).Return(nil)
			}

			service := NewBonusService(mockDB)
			bonus, err := service.CreateTransactionBonus(tt.userID, tt.amount, tt.transactionType)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, bonus)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, bonus)
				assert.Equal(t, tt.expectedAmount, bonus.Amount)
				assert.Equal(t, "transaction", bonus.Type)
				assert.Equal(t, "active", bonus.Status)

				// Проверяем, что срок действия установлен на 90 дней вперед
				expectedExpiry := time.Now().AddDate(0, 0, 90)
				assert.WithinDuration(t, expectedExpiry, bonus.ExpiresAt, 2*time.Second)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestBonusService_UseBonus(t *testing.T) {
	tests := []struct {
		name          string
		bonusID       string
		accountID     string
		setupMocks    func(*MockDatabase)
		expectedError bool
	}{
		{
			name:      "successful bonus usage",
			bonusID:   "bonus-1",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				bonus := &models.Bonus{
					ID:        "bonus-1",
					UserID:    "user-1",
					Type:      "welcome",
					Amount:    50.0,
					Status:    "active",
					ExpiresAt: time.Now().AddDate(0, 0, 30),
				}
				account := &models.Account{
					ID:       "account-1",
					UserID:   "user-1",
					Balance:  100.0,
					Currency: "USD",
				}

				mockDB.On("GetBonus", "bonus-1").Return(bonus, nil)
				mockDB.On("GetAccount", "account-1").Return(account, nil)
				mockDB.On("UpdateAccount", mock.AnythingOfType("*models.Account")).Return(nil)
				mockDB.On("UpdateBonus", mock.AnythingOfType("*models.Bonus")).Return(nil)
			},
			expectedError: false,
		},
		{
			name:      "bonus not active",
			bonusID:   "bonus-1",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				bonus := &models.Bonus{
					ID:        "bonus-1",
					UserID:    "user-1",
					Type:      "welcome",
					Amount:    50.0,
					Status:    "used",
					ExpiresAt: time.Now().AddDate(0, 0, 30),
				}

				mockDB.On("GetBonus", "bonus-1").Return(bonus, nil)
			},
			expectedError: true,
		},
		{
			name:      "bonus expired",
			bonusID:   "bonus-1",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				bonus := &models.Bonus{
					ID:        "bonus-1",
					UserID:    "user-1",
					Type:      "welcome",
					Amount:    50.0,
					Status:    "active",
					ExpiresAt: time.Now().AddDate(0, 0, -1), // expired yesterday
				}

				mockDB.On("GetBonus", "bonus-1").Return(bonus, nil)
				mockDB.On("UpdateBonus", mock.AnythingOfType("*models.Bonus")).Return(nil)
			},
			expectedError: true,
		},
		{
			name:      "wrong user account",
			bonusID:   "bonus-1",
			accountID: "account-1",
			setupMocks: func(mockDB *MockDatabase) {
				bonus := &models.Bonus{
					ID:        "bonus-1",
					UserID:    "user-1",
					Type:      "welcome",
					Amount:    50.0,
					Status:    "active",
					ExpiresAt: time.Now().AddDate(0, 0, 30),
				}
				account := &models.Account{
					ID:       "account-1",
					UserID:   "user-2", // different user
					Balance:  100.0,
					Currency: "USD",
				}

				mockDB.On("GetBonus", "bonus-1").Return(bonus, nil)
				mockDB.On("GetAccount", "account-1").Return(account, nil)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &MockDatabase{}
			tt.setupMocks(mockDB)

			service := NewBonusService(mockDB)
			err := service.UseBonus(tt.bonusID, tt.accountID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

func TestBonusService_GetActiveBonuses(t *testing.T) {
	mockDB := &MockDatabase{}
	bonuses := []*models.Bonus{
		{
			ID:        "bonus-1",
			UserID:    "user-1",
			Type:      "welcome",
			Amount:    50.0,
			Status:    "active",
			ExpiresAt: time.Now().AddDate(0, 0, 30),
		},
		{
			ID:        "bonus-2",
			UserID:    "user-1",
			Type:      "transaction",
			Amount:    25.0,
			Status:    "active",
			ExpiresAt: time.Now().AddDate(0, 0, 60),
		},
		{
			ID:        "bonus-3",
			UserID:    "user-1",
			Type:      "welcome",
			Amount:    100.0,
			Status:    "used", // неактивный
			ExpiresAt: time.Now().AddDate(0, 0, 30),
		},
		{
			ID:        "bonus-4",
			UserID:    "user-1",
			Type:      "welcome",
			Amount:    75.0,
			Status:    "active",
			ExpiresAt: time.Now().AddDate(0, 0, -1), // истекший
		},
	}

	mockDB.On("GetBonusesByUserID", "user-1").Return(bonuses, nil)

	service := NewBonusService(mockDB)
	result, err := service.GetActiveBonuses("user-1")

	assert.NoError(t, err)
	assert.Len(t, result, 2) // только активные и неистекшие бонусы

	// Проверяем, что возвращены только активные бонусы
	for _, bonus := range result {
		assert.Equal(t, "active", bonus.Status)
		assert.True(t, time.Now().Before(bonus.ExpiresAt))
	}

	mockDB.AssertExpectations(t)
}
