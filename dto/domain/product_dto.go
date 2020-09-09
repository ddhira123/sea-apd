package domain

import "github.com/williamchang80/sea-apd/domain/product"

type ProductDto struct {
	Product *product.Product `json:"product"`
}
