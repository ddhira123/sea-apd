package admin

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	admin2 "github.com/williamchang80/sea-apd/domain/admin"
	"github.com/williamchang80/sea-apd/dto/request/admin"
	"github.com/williamchang80/sea-apd/dto/response/base"

	"net/http"
)

type AdminController struct {
	usecase admin2.AdminUsecase
}

func NewAdminController(e *echo.Echo, a admin2.AdminUsecase) admin2.AdminController {
	c := &AdminController{
		usecase: a,
	}
	e.POST("api/user/admin", c.RegisterAdmin)
	return c
}

func (a *AdminController) RegisterAdmin(c echo.Context) error {
	var adminRequest admin.AdminRequest
	c.Bind(&adminRequest)

	if err := a.usecase.RegisterAdmin(adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}
