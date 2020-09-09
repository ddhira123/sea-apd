package auth

import (
	"github.com/williamchang80/sea-apd/domain/auth"
	user "github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/auth"
)

type AuthUsecase struct {
	repo user.UserRepository
}

func NewAuthUsecase(repository user.UserRepository) auth.AuthUsecase {
	return AuthUsecase{repo: repository}
}

func (a AuthUsecase) Login(request request.LoginRequest) (string, error) {
	panic("implement me")
}
