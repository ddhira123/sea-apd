package transfer

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/transfer"
	request "github.com/williamchang80/sea-apd/dto/request/transfer"
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

var (
	emptyCreateTransferHistoryRequest = request.CreateTransferHistoryRequest{}
)

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) GetTransferHistory(merchantId string) ([]transfer.Transfer, error) {
	if len(merchantId) == 0 {
		return nil, errors.New("mechantId cannot be empty")
	}
	return []transfer.Transfer{}, nil
}

func (m MockUsecase) CreateTransferHistory(request request.CreateTransferHistoryRequest) error {
	if request == emptyCreateTransferHistoryRequest{
		return errors.New("request cannot be empty")
	}
	return nil
}
