package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/merchant"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres/merchant"
	use_case "github.com/williamchang80/sea-apd/usecase/merchant"
)

type MerchantRoute struct {
	controller domain.MerchantController
	usecase    domain.MerchantUsecase
	repository domain.MerchantRepository
}

func NewMerchantRoute(e *echo.Echo) MerchantRoute {
	db := db.Postgres()
	if db != nil {
		d := db.AutoMigrate(&domain.Merchant{})
		d.AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	}
	repo := merchant.NewMerchantRepository(db)
	usecase := use_case.NewMerchantUsecase(repo)
	c := controller.NewMerchantController(e, usecase)
	return MerchantRoute{
		controller: c,
		usecase:    usecase,
		repository: repo,
	}
}
