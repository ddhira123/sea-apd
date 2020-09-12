package user

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

var (
	emptyRegisterUserRequest   = auth.RegisterUserRequest{}
	emptyUpdateUserRoleRequest = user2.UpdateUserRoleRequest{}
	emptyUpdateUserRequest     = user2.UpdateUserRequest{}
	mockUser                   = user.User{}
)

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) CreateUser(request auth.RegisterUserRequest) error {
	if request == emptyRegisterUserRequest {
		return errors.New("Cannot create user")
	}
	return nil
}

func (m MockUsecase) UpdateUserRole(request user2.UpdateUserRoleRequest) error {
	if request == emptyUpdateUserRoleRequest {
		return errors.New("Cannot update user role")
	}
	return nil
}

func (m MockUsecase) GetUserById(userId string) (*user.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("Cannot get user by id")
	}
	return &mockUser, nil
}

func (m MockUsecase) UpdateUser(request user2.UpdateUserRequest) error {
	if request == emptyUpdateUserRequest {
		return errors.New("Cannot update user")
	}
	return nil
}
