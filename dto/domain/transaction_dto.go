package domain

import "github.com/williamchang80/sea-apd/domain/transaction"

type TransactionDto struct {
	Transaction transaction.Transaction `json:"transaction"`
}