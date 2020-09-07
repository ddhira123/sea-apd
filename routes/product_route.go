package routes

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	product2 "github.com/williamchang80/sea-apd/repository/postgres/product"
	use_case "github.com/williamchang80/sea-apd/usecase/product"
)


func NewProductRoutes(e *echo.Echo) {
	db := db.Postgres()
	repo := product2.NewProductRepository(db)
	usecase := use_case.NewProductUseCase(repo)
	product.NewProductController(e, usecase)
}
