package request

import (
	"mime/multipart"
)

type Product struct {
	Name string
	Stock int
	Description string
	Price int
	Image multipart.File
}
