package user

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/user"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

var emptyUser = user.User{}
var emptyUserRequest = user2.UserRequest{}

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) GetUsers() ([]user.User, error) {
	return []user.User{}, nil
}

func (m MockUsecase) RegisterUser(request user2.UserRequest) error {
	if request == emptyUserRequest {
		return errors.New("Cannot Create User")
	}
	return nil
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}
