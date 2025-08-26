package api

import (
	"net/http"

	"petProjectMike/internal/config"
	"petProjectMike/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config             *config.Config
	transactionService *services.TransactionService
	bonusService       *services.BonusService
	accountService     *services.AccountService
	router             *gin.Engine
}

func NewServer(
	cfg *config.Config,
	transactionService *services.TransactionService,
	bonusService *services.BonusService,
	accountService *services.AccountService,
) *Server {
	server := &Server{
		config:             cfg,
		transactionService: transactionService,
		bonusService:       bonusService,
		accountService:     accountService,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	if s.config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.router = gin.Default()
	s.router.Use(gin.Logger(), gin.Recovery())

	s.router.GET("/health", s.healthCheck)

	v1 := s.router.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/:id", s.getAccount)
			accounts.GET("/:id/summary", s.getAccountSummary)
			accounts.GET("/user/:userID", s.getAccountsByUser)
			accounts.POST("/", s.createAccount)
			accounts.PUT("/:id", s.updateAccount)
			accounts.DELETE("/:id", s.deleteAccount)
		}

		transactions := v1.Group("/transactions")
		{
			transactions.GET("/:id", s.getTransaction)
			transactions.GET("/account/:accountID", s.getTransactionHistory)
			transactions.POST("/transfer", s.createTransfer)
			transactions.POST("/deposit", s.createDeposit)
			transactions.POST("/withdrawal", s.createWithdrawal)
		}

		bonuses := v1.Group("/bonuses")
		{
			bonuses.GET("/:id", s.getBonus)
			bonuses.GET("/user/:userID", s.getUserBonuses)
			bonuses.POST("/welcome", s.createWelcomeBonus)
			bonuses.POST("/use", s.useBonus)
		}

		users := v1.Group("/users")
		{
			users.GET("/:id", s.getUser)
			users.POST("/", s.createUser)
			users.PUT("/:id", s.updateUser)
			users.DELETE("/:id", s.deleteUser)
		}
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.config.Port)
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "banking-api", "version": "1.0.0"})
}
