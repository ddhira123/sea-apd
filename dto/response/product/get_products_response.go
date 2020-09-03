package product

import (
	"github.com/williamchang80/sea-apd/dto/domain"
)

type GetProductsResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data domain.ProductListDto `json:"data"`
}
