package domain

import "github.com/williamchang80/sea-apd/domain/product"

type ProductListDto struct {
	Products []product.Product `json:"products"`
}
