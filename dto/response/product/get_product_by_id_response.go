package product

import "github.com/williamchang80/sea-apd/dto/domain"

type GetProductByIdResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data domain.ProductDto `json:"data"`
}
