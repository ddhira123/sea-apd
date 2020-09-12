package admin

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/admin"
)

var (
	emptyAdmin        = user.User{}
	emptyAdminRequest = admin.AdminRequest{}
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) RegisterAdmin(req admin.AdminRequest) error {
	if req == emptyAdminRequest {
		return errors.New("Cannot register admin")
	}
	return nil
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}
