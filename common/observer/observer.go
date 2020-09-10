package observer

import (
	merchant3 "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionObserver interface {
	Update(transaction transaction.Transaction, m merchant3.MerchantUsecase) error
}
