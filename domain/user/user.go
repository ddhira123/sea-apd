package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/common/security"
	"github.com/williamchang80/sea-apd/domain"
	request "github.com/williamchang80/sea-apd/dto/request/user"
)

// User ...
type User struct {
	domain.Base
	Name     string `gorm:"size:50;not null;" json:"name"`
	Email    string `gorm:"unique;size:100;not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Role     string `gorm:"size:1;not null;" json:"role"`
	Phone    string `json:"phone"`
}

// UserRepository ...
type UserRepository interface {
	CreateUser(User) error
	GetUserByEmail(string) (*User, error)
	GetUsers() ([]User, error)
}

// UserUsecase ...
type UserUsecase interface {
	GetUsers() ([]User, error)
	RegisterUser(request.UserRequest) error
}

// UserController ...
type UserController interface {
	GetUsers(echo.Context) error
	RegisterUser(echo.Context) error
}

// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	scope.SetColumn("ID", id)

	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}
