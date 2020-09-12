package transaction

import (
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetTransactionHistoryResponse struct {
	base.BaseResponse
	Data domain.TransactionListDto `json:"data"`
}

type GetTransactionByIdResponse struct {
	base.BaseResponse
	Data domain.TransactionDto `json:"data"`
}

type GetCartItemsResponse struct {
	base.BaseResponse
	Data domain.CartDto `json:"data"`
}
