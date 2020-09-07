package transaction

import (
	"github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetTransactionHistoryResponse struct {
	base.BaseResponse
	Data []transaction.Transaction `json:"data"`
}

type GetTransactionByIdResponse struct {
	base.BaseResponse
	Data domain.TransactionDto `json:"data"`
}
