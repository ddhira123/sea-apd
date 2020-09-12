package product

import (
	"mime/multipart"
)

type ProductRequest struct {
	Name        string         `json:"name"`
	Stock       int            `json:"stock"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	MerchantId  string         `json:"merchant_id"`
	Image       multipart.File `json:"image"`
}
