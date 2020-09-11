package converter

import (
	transaction2 "github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

func MergePaymentRequestAndTransactionTotal(request transaction.PaymentRequest, transaction transaction2.Transaction,
	total int) transaction2.Transaction {
	return transaction2.Transaction{
		BankNumber:     request.BankNumber,
		BankName:       request.BankName,
		Amount:         total,
		CustomerId:     transaction.CustomerId,
		Status:         transaction.Status,
		MerchantId:     transaction.MerchantId,
		ProductDetails: transaction.ProductDetails,
	}
}
