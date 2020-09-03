package routes

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	use_case "github.com/williamchang80/sea-apd/usecase/product"
)


func NewProductRoutes(e *echo.Echo) {
	db := db.Postgres()
	repo := postgres.NewProductRepository(db)
	usecase := use_case.NewProductUseCase(repo)
	product.NewProductController(e, usecase)
}
