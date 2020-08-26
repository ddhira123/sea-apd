package product

type Product struct {
	Name        string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}

func NewProduct(name string, desc string, price int, image string, stock int) *Product {
	return &Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Image:       image,
		Stock:       stock,
	}
}
