package auth

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/dto/request/auth"
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

var (
	emptyAuthLoginRequest = auth.LoginRequest{
		Email:    "123@gmail.com",
		Password: "123",
	}
)

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}
func (m MockUsecase) Login(request auth.LoginRequest) (string, error) {
	if request == emptyAuthLoginRequest {
		return "", errors.New("cannot login")
	}
	return "token", nil
}
