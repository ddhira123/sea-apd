package user

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
	"github.com/williamchang80/sea-apd/dto/response/base"
	"net/http"
)

type UserController struct {
	usecase user.UserUsecase
}

func NewUserController(e *echo.Echo, uc user.UserUsecase) user.UserController {
	c := &UserController{
		usecase: uc,
	}
	e.POST("api/auth/register", c.CreateUser)
	e.PUT("api/user", c.UpdateUser)
	return c
}

func (u UserController) CreateUser(context echo.Context) error {
	var request auth.RegisterUserRequest
	context.Bind(&request)

	err := u.usecase.CreateUser(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return context.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (u UserController) UpdateUser(c echo.Context) error {
	var request user2.UpdateUserRequest
	c.Bind(&request)
	err := u.usecase.UpdateUser(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})

}
