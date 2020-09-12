package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/auth"
	domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	user2 "github.com/williamchang80/sea-apd/repository/postgres/user"
	"github.com/williamchang80/sea-apd/usecase/auth"
)

type AuthRoute struct {
	controller domain.AuthController
	usecase    domain.AuthUsecase
}

func NewAuthRoute(e *echo.Echo) AuthRoute {
	db := db.Postgres()
	repo :=  user2.NewUserRepository(db)
	usecase := auth.NewAuthUsecase(repo)
	c := controller.NewAuthController(e, usecase)
	return AuthRoute{
		controller: c,
		usecase:    usecase,
	}
}
