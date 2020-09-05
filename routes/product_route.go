package routes

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	domain "github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	product2 "github.com/williamchang80/sea-apd/repository/postgres/product"
	use_case "github.com/williamchang80/sea-apd/usecase/product"
)

type ProductRoute struct {
	controller domain.ProductController
	usecase    domain.ProductUsecase
	repository domain.ProductRepository
}

func NewProductRoutes(e *echo.Echo) ProductRoute {
	db := db.Postgres()
	repo := product2.NewProductRepository(db)
	usecase := use_case.NewProductUseCase(repo)
	controller := product.NewProductController(e, usecase)
	return ProductRoute{
		controller: controller,
		usecase:    usecase,
		repository: repo,
	}
}
