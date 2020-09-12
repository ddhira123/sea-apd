package admin

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/dto/request/admin"
)

type AdminUsecase interface {
	RegisterAdmin(request admin.AdminRequest) error
}

type AdminController interface {
	RegisterAdmin(echo.Context) error
}
