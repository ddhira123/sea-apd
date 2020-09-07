package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/user"
)

// UserUsecase ... definition
type UserUsecase struct {
	ur user.UserRepository
}

// ConvertToDomain ...
func ConvertToDomain(u request.UserRequest) user.User {
	return user.User{
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
		Phone: u.Phone,
		Role: u.Role,
	}
}

// NewUserUseCase is an usecase function for New user
func NewUserUseCase(u user.UserRepository) user.UserUsecase {
	return &UserUsecase{
		ur: u,
	}
}

// RegisterUser ...
func (s *UserUsecase) RegisterUser(user request.UserRequest) error {
	u, err := s.ur.GetUserByEmail(user.Email)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if u != nil {
		return errors.New("duplicate")
	}
	
	a := ConvertToDomain(user)
	err = s.ur.CreateUser(a)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserUsecase) GetUsers() ([]user.User, error) {
	u, err := s.ur.GetUsers()
	if err != nil {
		return nil, err
	}
	return u, nil
}