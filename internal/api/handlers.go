package api

import (
	"net/http"

	"petProjectMike/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := s.accountService.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (s *Server) getAccountSummary(c *gin.Context) {
	id := c.Param("id")
	summary, err := s.accountService.GetAccountSummary(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}

func (s *Server) getAccountsByUser(c *gin.Context) {
	userID := c.Param("userID")
	accounts, err := s.accountService.GetAccountsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (s *Server) createAccount(c *gin.Context) {
	var request struct {
		UserID   string `json:"user_id" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := s.accountService.CreateAccount(request.UserID, request.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (s *Server) updateAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account.ID = id
	if err := s.accountService.UpdateAccount(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (s *Server) deleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := s.accountService.DeleteAccount(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

func (s *Server) getTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction, err := s.transactionService.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (s *Server) getTransactionHistory(c *gin.Context) {
	accountID := c.Param("accountID")
	transactions, err := s.transactionService.GetTransactionHistory(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (s *Server) createTransfer(c *gin.Context) {
	var request struct {
		FromAccount string  `json:"from_account" binding:"required"`
		ToAccount   string  `json:"to_account" binding:"required"`
		Amount      float64 `json:"amount" binding:"required,gt=0"`
		Description string  `json:"description"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := s.transactionService.CreateTransfer(request.FromAccount, request.ToAccount, request.Amount, request.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

func (s *Server) createDeposit(c *gin.Context) {
	var request struct {
		AccountID   string  `json:"account_id" binding:"required"`
		Amount      float64 `json:"amount" binding:"required,gt=0"`
		Description string  `json:"description"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := s.transactionService.CreateDeposit(request.AccountID, request.Amount, request.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

func (s *Server) createWithdrawal(c *gin.Context) {
	var request struct {
		AccountID   string  `json:"account_id" binding:"required"`
		Amount      float64 `json:"amount" binding:"required,gt=0"`
		Description string  `json:"description"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := s.transactionService.CreateWithdrawal(request.AccountID, request.Amount, request.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

func (s *Server) getBonus(c *gin.Context) {
	id := c.Param("id")
	bonus, err := s.bonusService.GetBonus(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bonus)
}

func (s *Server) getUserBonuses(c *gin.Context) {
	userID := c.Param("userID")
	bonuses, err := s.bonusService.GetActiveBonuses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bonuses)
}

func (s *Server) createWelcomeBonus(c *gin.Context) {
	var request struct {
		UserID string  `json:"user_id" binding:"required"`
		Amount float64 `json:"amount" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bonus, err := s.bonusService.CreateWelcomeBonus(request.UserID, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bonus)
}

func (s *Server) useBonus(c *gin.Context) {
	var request struct {
		BonusID   string `json:"bonus_id" binding:"required"`
		AccountID string `json:"account_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.bonusService.UseBonus(request.BonusID, request.AccountID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bonus used successfully"})
}

func (s *Server) getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := s.accountService.GetDB().GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.accountService.GetDB().CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (s *Server) updateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := s.accountService.GetDB().UpdateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := s.accountService.GetDB().DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
