package user

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type MockUsecase struct {
	ctrl *gomock.Controller
}


func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) CreateUser(request auth.RegisterUserRequest) error {
	panic("implement me")
}

func (m MockUsecase) UpdateUserRole(request user2.UpdateUserRoleRequest) error {
	panic("implement me")
}