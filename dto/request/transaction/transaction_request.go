package transaction

import (
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
)

type TransactionRequest struct {
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	CustomerId string `json:"customer_id"`
	MerchantId string `json:"merchant_id"`
}

type UpdateTransactionRequest struct {
	TransactionId string                               `json:"transaction_id"`
	Status        transaction_status.TransactionStatus `json:"status"`
}

type PaymentRequest struct {
	CustomerId    string `json:"customer_id"`
	BankNumber    string `json:"bank_number"`
	BankName      string `json:"bank_name"`
	TransactionId string `json:"transaction_id"`
}

type CartRequest struct {
	ProductId     string `json:"product_id"`
	TransactionId string `json:"transaction_id"`
	Quantity      int    `json:"quantity"`
}

type CreateCartRequest struct {
	UserId     string `json:"user_id"`
	MerchantId string `json:"merchant_id"`
}
