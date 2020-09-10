package transaction

import "github.com/williamchang80/sea-apd/common/constants/transaction_status"

type TransactionRequest struct {
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
	Amount     int    `json:"amount"`
	UserId     string `json:"user_id"`
}

type UpdateTransactionRequest struct {
	TransactionId string                               `json:"transaction_id"`
	Status        transaction_status.TransactionStatus `json:"status"`
}
