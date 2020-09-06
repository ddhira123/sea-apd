package transaction

import (
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetTransactionByIdResponse struct {
	base.BaseResponse
	Data domain.TransactionDto `json:"data"`
}
