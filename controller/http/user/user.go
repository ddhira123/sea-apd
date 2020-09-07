package user

import (
	"fmt"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// mid "github.com/williamchang80/sea-apd/controller/middleware"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/domain"
	request "github.com/williamchang80/sea-apd/dto/request/user"
	"github.com/williamchang80/sea-apd/dto/response/base"
	response "github.com/williamchang80/sea-apd/dto/response/user"

	"net/http"
)

// UserController ...
type UserController struct {
	usecase user.UserUsecase
}

// NewUserController ...
func NewUserController(e *echo.Echo, uc user.UserUsecase) user.UserController {
	c := &UserController{
		usecase: uc,
	}
	e.POST("/register-user", c.RegisterUser)
	e.GET("/list-users", c.GetUsers)
	return c
}

// RegisterUser ...
func (a *UserController) RegisterUser(c echo.Context) error {
	var userRequest request.UserRequest
	c.Bind(&userRequest)
	fmt.Println(userRequest)

	if err := a.usecase.RegisterUser(userRequest); err != nil {
		if err.Error() == "duplicate" {
			return c.JSON(http.StatusBadRequest, &base.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Email has been taken",
			})
		}
		return c.JSON(http.StatusInternalServerError, &base.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Error",
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
	})
}

func (u *UserController) GetUsers(c echo.Context) error {
	users, err := u.usecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "User NOT FOUND",
		})
	}

	return c.JSON(http.StatusOK, &response.GetUsersResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
		Data:    domain.UserListDto{Users: users},
	})
}
