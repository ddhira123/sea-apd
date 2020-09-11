package user

import (
	"errors"
	auth2 "github.com/williamchang80/sea-apd/common/auth"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type UserUsecase struct {
	repo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) user.UserUsecase {
	return UserUsecase{repo: repo}
}

func convertRegisterRequestToUserDomain(request auth.RegisterUserRequest) user.User {
	return user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: auth2.HashPassword(request.Password),
		Role:     user_role.ToString(user_role.CUSTOMER),
	}
}

func (u UserUsecase) CreateUser(request auth.RegisterUserRequest) error {
	if request.Password != request.PasswordConfirmation {
		return errors.New("password and confirmation password must be same")
	}
	user := convertRegisterRequestToUserDomain(request)
	if err := u.repo.CreateUser(user); err != nil {
		return errors.New("email must be unique")
	}
	return nil
}

func (u UserUsecase) UpdateUserRole(request user2.UpdateUserRoleRequest) error {
	if err := u.repo.UpdateUserRole(user_role.ToString(request.Role), request.UserId); err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) GetUserById(userId string) (*user.User, error) {
	user, err := u.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
