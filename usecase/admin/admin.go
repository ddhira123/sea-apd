package user

import (
	"errors"
	admin2 "github.com/williamchang80/sea-apd/domain/admin"
	"os"

	"github.com/williamchang80/sea-apd/dto/request/auth"

	"github.com/williamchang80/sea-apd/common/constants/user_role"
	auth_domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
)

type AdminUsecase struct {
	ur      user.UserRepository
	usecase auth_domain.AuthUsecase
}

func NewAdminUseCase(a user.UserRepository, usecase auth_domain.AuthUsecase) admin2.AdminUsecase {
	return &AdminUsecase{
		ur:      a,
		usecase: usecase,
	}
}

func (s *AdminUsecase) RegisterAdmin(request admin.AdminRequest) error {
	authRequest := auth.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
	if _, err := s.usecase.Login(authRequest); err != nil {
		return err
	}

	u, err := s.ur.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("credential not match")
	}

	if !isValidToken(request.Token) {
		return errors.New("token not valid")
	}

	if err := s.ur.UpdateUserRole(user_role.ToString(user_role.ADMIN), u.ID); err != nil {
		return err
	}

	return nil
}

func isValidToken(token string) bool {
	if token == os.Getenv("ADMIN_TOKEN") {
		return true
	}

	return false
}
