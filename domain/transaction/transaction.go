package transaction

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

type Transaction struct {
	domain.Base
	Status     string `json:"status"`
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	User       user.User
	UserId     string `json:"user_id"`
}

type TransactionUsecase interface {
	CreateTransaction(transaction.TransactionRequest) error
	GetTransactionById(id string) (*Transaction, error)
	UpdateTransactionStatus(transaction.UpdateTransactionRequest) error
	GetTransactionHistory(userId string) ([]Transaction, error)
}

type TransactionController interface {
	CreateTransaction(echo.Context) error
	GetTransactionById(echo.Context) error
	UpdateTransactionStatus(echo.Context) error
	GetTransactionHistory(echo.Context) error
}

type TransactionRepository interface {
	CreateTransaction(Transaction) error
	GetTransactionById(string) (*Transaction, error)
	UpdateTransactionStatus(status string, id string) (*Transaction, error)
	GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]Transaction, error)
}
