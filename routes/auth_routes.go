package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/auth"
	domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/usecase/auth"
)

type AuthRoute struct {
	controller domain.AuthController
	usecase    domain.AuthUsecase
}

func NewAuthRoute(e *echo.Echo) AuthRoute {
	user := NewUserRoute(e)
	usecase := auth.NewAuthUsecase(user.repository)
	c := controller.NewAuthController(e, usecase)
	return AuthRoute{
		controller: c,
		usecase:    usecase,
	}
}
