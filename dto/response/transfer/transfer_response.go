package transfer

import (
	"github.com/williamchang80/sea-apd/domain/transfer"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetTransferResponse struct {
	base.BaseResponse
	Data []transfer.Transfer `json:"data"`
}
