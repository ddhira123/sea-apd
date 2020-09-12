package auth

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/dto/request/auth"
)

type AuthController interface {
	Login(echo.Context) error
}

type AuthUsecase interface {
	Login(request auth.LoginRequest) (string, error)
}
