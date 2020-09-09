package user

import (
	"errors"

	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/user"
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func (m MockRepository) GetUsers() ([]domain.User, error) {
	m.ctrl.T.Helper()
	return []domain.User{}, nil
}

func (m MockRepository) GetUserByEmail(email string) (*domain.User, error) {
	if email != "" {
		return &domain.User{
			Name:     "Mock Name",
			Email:    "Mock Email",
			Password: "Mock Password",
			Phone:    "Mock Phone",
			Role:     "4",
		}, nil
	}
	return nil, errors.New("Cannot Get User By Email")
}

func (m MockRepository) CreateUser(user domain.User) error {
	var p = domain.User{}
	if user == p {
		return errors.New("Cannot Create User")
	}
	return nil
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}
