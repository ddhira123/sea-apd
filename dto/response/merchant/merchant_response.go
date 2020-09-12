package merchant

import (
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetMerchantBalanceResponse struct {
	base.BaseResponse
	Data domain.MerchantBalanceDto `json:"data"`
}

type GetMerchantByIdResponse struct {
	base.BaseResponse
	Data domain.MerchantDto `json:"data"`
}

type GetMerchantsResponse struct {
	base.BaseResponse
	Data domain.MerchantListDto `json:"data"`
}
