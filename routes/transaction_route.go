package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/transaction"
	domain "github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres/transaction"
	usecase "github.com/williamchang80/sea-apd/usecase/transaction"
)

type TransactionRoute struct {
	controller domain.TransactionController
	usecase    domain.TransactionUsecase
	repository domain.TransactionRepository
}

func NewTransactionRoute(e *echo.Echo) Routes {
	merchantRoute := NewMerchantRoute(e)
	productRoute := NewProductRoutes(e)
	db := db.Postgres()
	if db != nil {
		d := db.AutoMigrate(&domain.Transaction{}, &domain.ProductTransaction{})
		d.Model(&domain.Transaction{}).AddForeignKey("customer_id", "users(id)",
			"CASCADE", "CASCADE")
		d.Model(&domain.Transaction{}).AddForeignKey("merchant_id", "merchants(id)",
			"CASCADE", "CASCADE")
		d.Model(&domain.ProductTransaction{}).AddForeignKey("product_id", "products(id)",
			"CASCADE", "CASCADE")
		d.Model(&domain.ProductTransaction{}).AddForeignKey("transaction_id", "transactions(id)",
			"CASCADE", "CASCADE")
	}
	repo := transaction.NewTransactionRepository(db)
	u := usecase.NewTransactionUsecase(repo, merchantRoute.Usecase, productRoute.Usecase)
	controller := controller.NewTransactionController(e, u)
	return Routes{
		Controller: controller,
		Usecase:    u,
		Repository: repo,
	}
}
