package product

import (
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetProductsResponse struct {
	base.BaseResponse
	Data domain.ProductListDto `json:"data"`
}
