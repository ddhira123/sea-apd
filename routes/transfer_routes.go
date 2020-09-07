package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/transfer"
	domain "github.com/williamchang80/sea-apd/domain/transfer"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	repository "github.com/williamchang80/sea-apd/repository/postgres/transfer"
	"github.com/williamchang80/sea-apd/usecase/transfer"
)

type TransferRoute struct {
	controller domain.TransferController
	usecase    domain.TransferUsecase
	repository domain.TransferRepository
}

func NewTransferRoute(e *echo.Echo) TransferRoute {
	db := db.Postgres()
	merchant := NewMerchantRoute(e)
	if db != nil {
		d := db.AutoMigrate(&domain.Transfer{})
		d.AddForeignKey("merchant_id", "merchants(id)", "CASCADE", "CASCADE")
	}
	repo := repository.NewTransferRepository(db)
	usecase := transfer.NewTransferUsecase(repo, merchant.usecase)
	c := controller.NewTransferController(e, usecase)
	return TransferRoute{
		controller: c,
		usecase:    usecase,
		repository: repo,
	}

}
