package transaction

import "github.com/williamchang80/sea-apd/domain/transaction"

type GetTransactionHistoryResponse struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    []transaction.Transaction `json:"data"`
}
