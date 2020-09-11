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
	Usecase    domain.MerchantUsecase
	repository domain.MerchantRepository
}

func NewMerchantRoute(e *echo.Echo) MerchantRoute {
	db := db.Postgres()
	userRoute := NewUserRoute(e)
	if db != nil {
		d := db.AutoMigrate(&domain.Merchant{})
		d.AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	}
	repo := merchant.NewMerchantRepository(db)
	u := use_case.NewMerchantUsecase(repo, userRoute.usecase)
	c := controller.NewMerchantController(e, u)
	return MerchantRoute{
		controller: c,
		Usecase:    u,
		repository: repo,
	}
}
