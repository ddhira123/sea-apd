package user

import (
	"errors"
	auth2 "github.com/williamchang80/sea-apd/common/auth"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/domain"
	auth_domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type UserUsecase struct {
	repo    user.UserRepository
	usecase auth_domain.AuthUsecase
}

func NewUserUsecase(repo user.UserRepository, usecase auth_domain.AuthUsecase) user.UserUsecase {
	return UserUsecase{repo: repo, usecase: usecase}
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

func (u UserUsecase) UpdateUser(request user2.UpdateUserRequest) error {
	authRequest := auth.LoginRequest{
		Email:    request.OldEmail,
		Password: request.OldPassword,
	}
	if _, err := u.usecase.Login(authRequest); err != nil {
		return err
	}
	us, err := u.GetUserById(request.UserId)
	if us.Email != request.OldEmail || err != nil {
		return errors.New("credential doesnt match")
	}

	user := getUpdatedUserLoginInformation(request, us.ID)
	if err := u.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func getUpdatedUserLoginInformation(request user2.UpdateUserRequest,
	userId string) user.User {
	email := request.OldEmail
	if request.NewEmail != "" {
		email = request.NewEmail
	}
	u := user.User{
		Base: domain.Base{
			ID: userId,
		},
		Email: email,
	}
	if request.NewPassword != "" {
		u.Password = auth2.HashPassword(request.NewPassword)
	}

	return u
}
