package auth

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/auth"
	request "github.com/williamchang80/sea-apd/dto/request/auth"
	auth_response "github.com/williamchang80/sea-apd/dto/response/auth"
	"github.com/williamchang80/sea-apd/dto/response/base"
	"net/http"
)

type AuthController struct {
	usecase auth.AuthUsecase
}

func NewAuthController(echo *echo.Echo, a auth.AuthUsecase) auth.AuthController {
	c := AuthController{usecase: a}
	echo.POST("api/auth/login", c.Login)
	return c
}

func (a AuthController) Login(context echo.Context) error {
	var loginRequest request.LoginRequest
	context.Bind(&loginRequest)
	authToken, err := a.usecase.Login(loginRequest)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, base.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: response.UNAUTHENTICED,
		})
	}
	return context.JSON(http.StatusOK, auth_response.LoginResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: response.SUCCESS,
		},
		Token: authToken,
	})
}
