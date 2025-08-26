package database

import (
	"errors"
	"sync"

	"petProjectMike/internal/models"
)

type InMemoryDB struct {
	accounts     map[string]*models.Account
	transactions map[string]*models.Transaction
	bonuses      map[string]*models.Bonus
	users        map[string]*models.User
	mutex        sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	db := &InMemoryDB{
		accounts:     make(map[string]*models.Account),
		transactions: make(map[string]*models.Transaction),
		bonuses:      make(map[string]*models.Bonus),
		users:        make(map[string]*models.User),
	}
	db.seedData()
	return db
}

func (db *InMemoryDB) seedData() {
	testUser := &models.User{ID: "user-1", Email: "test@example.com", Name: "Test User"}
	db.users[testUser.ID] = testUser

	testAccount := &models.Account{ID: "account-1", UserID: testUser.ID, Balance: 1000.0, Currency: "USD"}
	db.accounts[testAccount.ID] = testAccount

	testBonus := &models.Bonus{ID: "bonus-1", UserID: testUser.ID, Type: "welcome", Amount: 50.0, Status: "active"}
	db.bonuses[testBonus.ID] = testBonus
}

// Account
func (db *InMemoryDB) CreateAccount(account *models.Account) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.accounts[account.ID]; exists {
		return errors.New("account already exists")
	}
	db.accounts[account.ID] = account
	return nil
}

func (db *InMemoryDB) GetAccount(id string) (*models.Account, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	account, exists := db.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func (db *InMemoryDB) GetAccountsByUserID(userID string) ([]*models.Account, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var accounts []*models.Account
	for _, account := range db.accounts {
		if account.UserID == userID {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}

func (db *InMemoryDB) UpdateAccount(account *models.Account) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.accounts[account.ID]; !exists {
		return errors.New("account not found")
	}
	db.accounts[account.ID] = account
	return nil
}

func (db *InMemoryDB) DeleteAccount(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.accounts[id]; !exists {
		return errors.New("account not found")
	}
	delete(db.accounts, id)
	return nil
}

// Transaction
func (db *InMemoryDB) CreateTransaction(transaction *models.Transaction) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.transactions[transaction.ID]; exists {
		return errors.New("transaction already exists")
	}
	db.transactions[transaction.ID] = transaction
	return nil
}

func (db *InMemoryDB) GetTransaction(id string) (*models.Transaction, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	transaction, exists := db.transactions[id]
	if !exists {
		return nil, errors.New("transaction not found")
	}
	return transaction, nil
}

func (db *InMemoryDB) GetTransactionsByAccount(accountID string) ([]*models.Transaction, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var transactions []*models.Transaction
	for _, transaction := range db.transactions {
		if transaction.FromAccount == accountID || transaction.ToAccount == accountID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}

func (db *InMemoryDB) UpdateTransaction(transaction *models.Transaction) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.transactions[transaction.ID]; !exists {
		return errors.New("transaction not found")
	}
	db.transactions[transaction.ID] = transaction
	return nil
}

func (db *InMemoryDB) DeleteTransaction(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.transactions[id]; !exists {
		return errors.New("transaction not found")
	}
	delete(db.transactions, id)
	return nil
}

// Bonus
func (db *InMemoryDB) CreateBonus(bonus *models.Bonus) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.bonuses[bonus.ID]; exists {
		return errors.New("bonus already exists")
	}
	db.bonuses[bonus.ID] = bonus
	return nil
}

func (db *InMemoryDB) GetBonus(id string) (*models.Bonus, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	bonus, exists := db.bonuses[id]
	if !exists {
		return nil, errors.New("bonus not found")
	}
	return bonus, nil
}

func (db *InMemoryDB) GetBonusesByUserID(userID string) ([]*models.Bonus, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var bonuses []*models.Bonus
	for _, bonus := range db.bonuses {
		if bonus.UserID == userID {
			bonuses = append(bonuses, bonus)
		}
	}
	return bonuses, nil
}

func (db *InMemoryDB) UpdateBonus(bonus *models.Bonus) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.bonuses[bonus.ID]; !exists {
		return errors.New("bonus not found")
	}
	db.bonuses[bonus.ID] = bonus
	return nil
}

func (db *InMemoryDB) DeleteBonus(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.bonuses[id]; !exists {
		return errors.New("bonus not found")
	}
	delete(db.bonuses, id)
	return nil
}

// User
func (db *InMemoryDB) CreateUser(user *models.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	db.users[user.ID] = user
	return nil
}

func (db *InMemoryDB) GetUser(id string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	user, exists := db.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (db *InMemoryDB) GetUserByEmail(email string) (*models.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	for _, user := range db.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (db *InMemoryDB) UpdateUser(user *models.User) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	db.users[user.ID] = user
	return nil
}

func (db *InMemoryDB) DeleteUser(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(db.users, id)
	return nil
}
