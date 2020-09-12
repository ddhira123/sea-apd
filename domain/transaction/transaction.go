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
								      AssociationForeignKey:TransactionId"`
}

type ProductTransaction struct {
	ProductId     string    `json:"product_id"`
	TransactionId string    `json:"transaction_id"`
	Quantity      int       `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransactionUsecase interface {
	GetTransactionById(id string) (*Transaction, error)
	UpdateTransactionStatus(transaction.UpdateTransactionRequest) error
	GetTransactionHistory(userId string) ([]Transaction, error)
	GetMerchantRequestItem(merchantId string) ([]Transaction, error)
	PayTransaction(request transaction.PaymentRequest) error
	AddCartItem(request transaction.CartRequest) error
	RemoveCartItem(request transaction.CartRequest) error
	UpdateCartItem(request transaction.CartRequest) error
	GetCartItems(id string) ([]ProductTransaction, error)
	CreateCart(request transaction.CreateCartRequest) error
}

type TransactionController interface {
	GetTransactionById(echo.Context) error
	UpdateTransactionStatus(echo.Context) error
	GetTransactionHistory(echo.Context) error
	GetMerchantRequestItem(echo.Context) error
	PayTransaction(echo.Context) error
	AddCartItem(echo.Context) error
	RemoveCartItem(echo.Context) error
	UpdateCartItem(echo.Context) error
	GetCartItems(echo.Context) error
	CreateCart(echo.Context) error
}

type TransactionRepository interface {
	GetTransactionById(string) (*Transaction, error)
	UpdateTransactionStatus(status string, id string) (*Transaction, error)
	GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]Transaction, error)
	GetMerchantRequestItem(merchantId string) ([]Transaction, error)
	UpdateTransaction(transaction Transaction) error
	AddCartItem(cart ProductTransaction) error
	RemoveCartItem(cart ProductTransaction) error
	UpdateCartItem(cart ProductTransaction) error
	GetCartItems(id string) ([]ProductTransaction, error)
	CreateCart(transaction Transaction) error
}
