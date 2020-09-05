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
	db := db.Postgres()
	repo := transaction.NewTransactionRepository(db)
	usecase := usecase.NewTransactionUsecase(repo, merchantRoute.usecase)
	controller := controller.NewTransactionController(e, usecase)
	return Routes{
		controller: controller,
		usecase:    usecase,
		repository: repo,
	}
}
