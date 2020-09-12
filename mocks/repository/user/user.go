package user

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/user"
)

var (
	emptyUser = user.User{}
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) CreateUser(user user.User) error {
	if user == emptyUser {
		return errors.New("cannot create user")
	}
	return nil
}

func (m MockRepository) GetUserByEmail(email string) (*user.User, error) {
	if len(email) == 0 {
		return nil, errors.New("cannot get user by email")
	}
	return &emptyUser, nil
}

func (m MockRepository) UpdateUserRole(role string, userId string) error {
	if len(role) == 0 || len(userId) == 0 {
		return errors.New("cannot update user role")
	}
	return nil
}

func (m MockRepository) GetUserById(userId string) (*user.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("cannot get user by id")
	}
	return &emptyUser, nil
}

func (m MockRepository) UpdateUser(user user.User) error {
	if user == emptyUser {
		return errors.New("cannot update user")
	}
	return nil
}
