package transaction

import (
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	domain "github.com/williamchang80/sea-apd/domain/transaction"
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
	CustomerId string `json:"customer_id"`
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	TransactionId string `json:"transaction_id"`
}

func MergePaymentRequestAndTransactionTotal(request PaymentRequest, transaction domain.Transaction,
	total int) domain.Transaction {
	return domain.Transaction{
		BankNumber:     request.BankNumber,
		BankName:       request.BankName,
		Amount:         total,
		CustomerId:     transaction.CustomerId,
		Status:         transaction.Status,
		MerchantId:     transaction.MerchantId,
		ProductDetails: transaction.ProductDetails,
	}
}
