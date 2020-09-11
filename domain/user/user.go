package user

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	"github.com/williamchang80/sea-apd/dto/request/user"
)

type User struct {
	domain.Base
	Name     string `gorm:"size:50;not null;" json:"name"`
	Email    string `gorm:"unique;unique;size:100;not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Role     string `gorm:"not null;" json:"role"`
}

type UserRepository interface {
	CreateUser(User) error
	GetUserByEmail(email string) (*User, error)
	UpdateUserRole(role string, userId string) error
	GetUserById(userId string) (*User, error)
	UpdateUser(User) error
}

type UserUsecase interface {
	CreateUser(request auth.RegisterUserRequest) error
	UpdateUserRole(request user.UpdateUserRoleRequest) error
	GetUserById(userId string) (*User, error)
	UpdateUser(request user.UpdateUserRequest) error
}

type UserController interface {
	CreateUser(echo.Context) error
	UpdateUser(echo.Context) error
}
