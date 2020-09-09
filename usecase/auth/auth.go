package auth

import (
	"errors"
	"github.com/williamchang80/sea-apd/common/auth"
	auth2 "github.com/williamchang80/sea-apd/domain/auth"
	user "github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/auth"
)

type AuthUsecase struct {
	repo user.UserRepository
}

func NewAuthUsecase(repository user.UserRepository) auth2.AuthUsecase {
	return AuthUsecase{repo: repository}
}

func (a AuthUsecase) Login(request request.LoginRequest) (string, error) {
	user, err := a.repo.GetUserByEmail(request.Email)
	if err != nil || !auth.IsMatchedPassword(user.Password, request.Password) {
		return "", errors.New("Password and email not matched")
	}
	token, err := auth.GenerateToken(user)
	if err != nil {
		return "", errors.New("Password and email not matched")
	}
	return token, nil
}
