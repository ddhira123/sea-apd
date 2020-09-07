package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/transaction"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres/transaction"
	usecase "github.com/williamchang80/sea-apd/usecase/transaction"
)

func NewTransactionRoute(e *echo.Echo) {
	d := db.Postgres()
	repo := transaction.NewTransactionRepository(d)
	u := usecase.NewTransactionUsecase(repo)
	controller.NewTransactionController(e, u)
}
