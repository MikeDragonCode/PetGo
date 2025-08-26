package database

import (
	"fmt"
	"testing"
	"time"

	"petProjectMike/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryDB(t *testing.T) {
	db := NewInMemoryDB()

	assert.NotNil(t, db)

	// Проверяем, что тестовые данные созданы
	user, err := db.GetUser("user-1")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)

	account, err := db.GetAccount("account-1")
	assert.NoError(t, err)
	assert.Equal(t, 1000.0, account.Balance)

	bonus, err := db.GetBonus("bonus-1")
	assert.NoError(t, err)
	assert.Equal(t, "welcome", bonus.Type)
}

func TestInMemoryDB_CreateAndGetAccount(t *testing.T) {
	db := NewInMemoryDB()

	account := &models.Account{
		ID:       "test-account",
		UserID:   "user-1",
		Balance:  500.0,
		Currency: "EUR",
	}

	// Создаем счет
	err := db.CreateAccount(account)
	assert.NoError(t, err)

	// Получаем счет
	retrievedAccount, err := db.GetAccount("test-account")
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccount)
}

func TestInMemoryDB_CreateDuplicateAccount(t *testing.T) {
	db := NewInMemoryDB()

	account := &models.Account{
		ID:       "duplicate-account",
		UserID:   "user-1",
		Balance:  500.0,
		Currency: "EUR",
	}

	// Создаем счет первый раз
	err := db.CreateAccount(account)
	assert.NoError(t, err)

	// Пытаемся создать дубликат
	err = db.CreateAccount(account)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestInMemoryDB_GetNonExistentAccount(t *testing.T) {
	db := NewInMemoryDB()

	account, err := db.GetAccount("non-existent")
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.Contains(t, err.Error(), "not found")
}

func TestInMemoryDB_UpdateAccount(t *testing.T) {
	db := NewInMemoryDB()

	// Создаем счет
	account := &models.Account{
		ID:       "update-test",
		UserID:   "user-1",
		Balance:  100.0,
		Currency: "USD",
	}
	err := db.CreateAccount(account)
	assert.NoError(t, err)

	// Обновляем счет
	account.Balance = 200.0
	err = db.UpdateAccount(account)
	assert.NoError(t, err)

	// Проверяем обновление
	updatedAccount, err := db.GetAccount("update-test")
	assert.NoError(t, err)
	assert.Equal(t, 200.0, updatedAccount.Balance)
}

func TestInMemoryDB_DeleteAccount(t *testing.T) {
	db := NewInMemoryDB()

	// Создаем счет
	account := &models.Account{
		ID:       "delete-test",
		UserID:   "user-1",
		Balance:  0.0, // Пустой счет
		Currency: "USD",
	}
	err := db.CreateAccount(account)
	assert.NoError(t, err)

	// Удаляем счет
	err = db.DeleteAccount("delete-test")
	assert.NoError(t, err)

	// Проверяем, что счет удален
	_, err = db.GetAccount("delete-test")
	assert.Error(t, err)
}

func TestInMemoryDB_CreateAndGetTransaction(t *testing.T) {
	db := NewInMemoryDB()

	transaction := &models.Transaction{
		ID:          "test-transaction",
		FromAccount: "account-1",
		ToAccount:   "account-2",
		Amount:      100.0,
		Type:        "transfer",
		Status:      "pending",
		Description: "test transfer",
	}

	// Создаем транзакцию
	err := db.CreateTransaction(transaction)
	assert.NoError(t, err)

	// Получаем транзакцию
	retrievedTransaction, err := db.GetTransaction("test-transaction")
	assert.NoError(t, err)
	assert.Equal(t, transaction, retrievedTransaction)
}

func TestInMemoryDB_GetTransactionsByAccount(t *testing.T) {
	db := NewInMemoryDB()

	// Создаем несколько транзакций
	transaction1 := &models.Transaction{
		ID:          "txn-1",
		FromAccount: "account-1",
		ToAccount:   "account-2",
		Amount:      100.0,
		Type:        "transfer",
		Status:      "completed",
	}

	transaction2 := &models.Transaction{
		ID:          "txn-2",
		FromAccount: "account-3",
		ToAccount:   "account-1",
		Amount:      50.0,
		Type:        "transfer",
		Status:      "completed",
	}

	err := db.CreateTransaction(transaction1)
	assert.NoError(t, err)

	err = db.CreateTransaction(transaction2)
	assert.NoError(t, err)

	// Получаем транзакции для account-1
	transactions, err := db.GetTransactionsByAccount("account-1")
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)
}

func TestInMemoryDB_CreateAndGetBonus(t *testing.T) {
	db := NewInMemoryDB()

	bonus := &models.Bonus{
		ID:        "test-bonus",
		UserID:    "user-1",
		Type:      "referral",
		Amount:    25.0,
		Status:    "active",
		ExpiresAt: time.Now().AddDate(0, 0, 30),
	}

	// Создаем бонус
	err := db.CreateBonus(bonus)
	assert.NoError(t, err)

	// Получаем бонус
	retrievedBonus, err := db.GetBonus("test-bonus")
	assert.NoError(t, err)
	assert.Equal(t, bonus, retrievedBonus)
}

func TestInMemoryDB_GetBonusesByUserID(t *testing.T) {
	db := NewInMemoryDB()

	// Создаем несколько бонусов для одного пользователя с уникальными ID
	bonus1 := &models.Bonus{
		ID:        "bonus-u2-1",
		UserID:    "user-2",
		Type:      "welcome",
		Amount:    50.0,
		Status:    "active",
		ExpiresAt: time.Now().AddDate(0, 0, 30),
	}

	bonus2 := &models.Bonus{
		ID:        "bonus-u2-2",
		UserID:    "user-2",
		Type:      "transaction",
		Amount:    25.0,
		Status:    "active",
		ExpiresAt: time.Now().AddDate(0, 0, 90),
	}

	err := db.CreateBonus(bonus1)
	assert.NoError(t, err)

	err = db.CreateBonus(bonus2)
	assert.NoError(t, err)

	// Получаем бонусы для user-2
	bonuses, err := db.GetBonusesByUserID("user-2")
	assert.NoError(t, err)
	// Может присутствовать сидовый bonus-1 для user-1, но нас интересуют только user-2
	assert.Len(t, bonuses, 2)
}

func TestInMemoryDB_CreateAndGetUser(t *testing.T) {
	db := NewInMemoryDB()

	user := &models.User{
		ID:    "test-user",
		Email: "testuser@example.com",
		Name:  "Test User",
	}

	// Создаем пользователя
	err := db.CreateUser(user)
	assert.NoError(t, err)

	// Получаем пользователя
	retrievedUser, err := db.GetUser("test-user")
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)
}

func TestInMemoryDB_GetUserByEmail(t *testing.T) {
	db := NewInMemoryDB()

	// Создаем пользователя
	user := &models.User{
		ID:    "email-user",
		Email: "emailuser@example.com",
		Name:  "Email User",
	}

	err := db.CreateUser(user)
	assert.NoError(t, err)

	// Ищем пользователя по email
	foundUser, err := db.GetUserByEmail("emailuser@example.com")
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	// Ищем несуществующего пользователя
	_, err = db.GetUserByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestInMemoryDB_ConcurrentAccess(t *testing.T) {
	db := NewInMemoryDB()

	// Тестируем конкурентный доступ
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			account := &models.Account{
				ID:       fmt.Sprintf("concurrent-%d", id),
				UserID:   "user-1",
				Balance:  float64(id * 100),
				Currency: "USD",
			}

			err := db.CreateAccount(account)
			assert.NoError(t, err)

			// Читаем созданный счет
			retrievedAccount, err := db.GetAccount(account.ID)
			assert.NoError(t, err)
			assert.Equal(t, account, retrievedAccount)

			done <- true
		}(i)
	}

	// Ждем завершения всех горутин
	for i := 0; i < 10; i++ {
		<-done
	}

	// Проверяем, что все счета созданы
	for i := 0; i < 10; i++ {
		account, err := db.GetAccount(fmt.Sprintf("concurrent-%d", i))
		assert.NoError(t, err)
		assert.Equal(t, float64(i*100), account.Balance)
	}
}
