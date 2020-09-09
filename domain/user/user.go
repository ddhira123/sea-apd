package user

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/auth"
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
}

type UserUsecase interface {
	CreateUser(request auth.RegisterUserRequest) error
}

type UserController interface {
	CreateUser(echo.Context) error
}
