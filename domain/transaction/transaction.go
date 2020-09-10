package transaction

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
	"time"
)

type Transaction struct {
	domain.Base
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	CustomerId string `json:"customer_id"`
	Status     string `json:"status"`
	MerchantId string `json:"merchant_id"`
	ProductDetails []ProductTransaction `json:"product_details" gorm:"many2many:product_transactions;
								      AssociationForeignKey:TransactionId""`
}

type ProductTransaction struct {
	ProductId     string    `json:"product_id"`
	TransactionId string    `json:"transaction_id"`
	Quantity      int       `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransactionUsecase interface {
	CreateTransaction(transaction.TransactionRequest) error
	GetTransactionById(id string) (*Transaction, error)
	UpdateTransactionStatus(transaction.UpdateTransactionRequest) error
	GetTransactionHistory(userId string) ([]Transaction, error)
	GetMerchantRequestItem(merchantId string) ([]Transaction, error)
}

type TransactionController interface {
	CreateTransaction(echo.Context) error
	GetTransactionById(echo.Context) error
	UpdateTransactionStatus(echo.Context) error
	GetTransactionHistory(echo.Context) error
	GetMerchantRequestItem(echo.Context) error
}

type TransactionRepository interface {
	CreateTransaction(Transaction) error
	GetTransactionById(string) (*Transaction, error)
	UpdateTransactionStatus(status string, id string) (*Transaction, error)
	GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]Transaction, error)
	GetMerchantRequestItem(merchantId string) ([]Transaction, error)
}
