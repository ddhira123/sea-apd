package product

import (
	"mime/multipart"
)

type ProductRequest struct {
	Name string
	Stock int
	Description string
	Price int
	Image multipart.File
}
