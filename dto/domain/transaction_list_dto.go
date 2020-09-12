package domain

import (
	"github.com/williamchang80/sea-apd/domain/transaction"
)

type TransactionListDto struct {
	Transactions []transaction.Transaction `json:"transactions"`
}

