package converter

import (
	"github.com/williamchang80/sea-apd/domain"
	transaction2 "github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

func MergePaymentRequestAndTransactionTotal(request transaction.PaymentRequest, transaction transaction2.Transaction,
	total int) transaction2.Transaction {
	return transaction2.Transaction{
		Base: domain.Base{
			ID:        transaction.ID,
		},
		BankNumber:     request.BankNumber,
		BankName:       request.BankName,
		Amount:         total,
		CustomerId:     transaction.CustomerId,
		Status:         transaction.Status,
		MerchantId:     transaction.MerchantId,
		ProductDetails: transaction.ProductDetails,
	}
}
