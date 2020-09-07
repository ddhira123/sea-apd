package transfer

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/merchant"
	domain "github.com/williamchang80/sea-apd/domain/transfer"
	"github.com/williamchang80/sea-apd/dto/request/transfer"
	request "github.com/williamchang80/sea-apd/dto/request/transfer"
	transfer2 "github.com/williamchang80/sea-apd/mocks/repository/transfer"
	merchant2 "github.com/williamchang80/sea-apd/mocks/usecase/merchant"
	"reflect"
	"testing"
)

var (
	mockId                       = "1"
	mockUpdateTransactionRequest = transfer.CreateTransferHistoryRequest{
		BankNumber: "1",
		BankName:   "name",
		Amount:     100,
		MerchantId: "1",
	}
	mockTransferEntity = domain.Transfer{
		Amount:     100,
		BankName:   "name",
		BankNumber: "1",
		MerchantId: "1",
	}
)

func TestNewTransferUsecase(t *testing.T) {
	type args struct {
		repository domain.TransferRepository
		usecase    merchant.MerchantUsecase
	}
	tests := []struct {
		name string
		args args
		want domain.TransferUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
				usecase:    nil,
			},
			want: &TransferUsecase{
				repo:            nil,
				merchantUsecase: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransferUsecase(tt.args.repository, tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransferUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToDomain(t *testing.T) {
	type args struct {
		productRequest request.CreateTransferHistoryRequest
	}
	tests := []struct {
		name string
		args args
		want domain.Transfer
	}{
		{
			name: "success",
			args: args{
				productRequest: mockUpdateTransactionRequest,
			},
			want: mockTransferEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCreateTransferRequestToDomain(tt.args.productRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertCreateTransferRequestToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferUsecase_GetTransferHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		want     []domain.Transfer
		initMock func() domain.TransferUsecase
	}{
		{
			name:    "success",
			wantErr: false,
			args: args{
				merchantId: "1",
			},
			want: []domain.Transfer{},
			initMock: func() domain.TransferUsecase {
				t := transfer2.NewMockRepository(ctrl)
				return NewTransferUsecase(t, nil)
			},
		},
		{
			name:    "failed with empty request",
			wantErr: true,
			args: args{
				merchantId: "",
			},
			initMock: func() domain.TransferUsecase {
				t := transfer2.NewMockRepository(ctrl)
				return NewTransferUsecase(t, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			got, err := c.GetTransferHistory(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("TransferUsecase.GetTransferHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferUsecase.GetTransferHistory() Find() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestTransferUsecase_CreateTransferHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.CreateTransferHistoryRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.TransferUsecase
	}{
		{
			name:    "success",
			wantErr: false,
			args: args{
				request: mockUpdateTransactionRequest,
			},
			initMock: func() domain.TransferUsecase {
				t := transfer2.NewMockRepository(ctrl)
				u := merchant2.NewMockUsecase(ctrl)
				return NewTransferUsecase(t, u)
			},
		},
		{
			name:    "failed with empty request",
			wantErr: true,
			args: args{
				request: request.CreateTransferHistoryRequest{},
			},
			initMock: func() domain.TransferUsecase {
				t := transfer2.NewMockRepository(ctrl)
				u := merchant2.NewMockUsecase(ctrl)
				return NewTransferUsecase(t, u)
			},
		},
		{
			name:    "failed with more amount than balance",
			wantErr: true,
			args: args{
				request: request.CreateTransferHistoryRequest{
					BankNumber: "test",
					BankName:   "test",
					Amount:     -1000000,
					MerchantId: "1",
				},
			},
			initMock: func() domain.TransferUsecase {
				t := transfer2.NewMockRepository(ctrl)
				u := merchant2.NewMockUsecase(ctrl)
				return NewTransferUsecase(t, u)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.CreateTransferHistory(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("TransferUsecase.CreateTransferHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
