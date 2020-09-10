package observer

import (
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionObserver interface {
	Update(transaction transaction.Transaction, usecase transaction.TransactionUsecase) error
}
