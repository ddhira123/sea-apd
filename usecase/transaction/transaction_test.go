package transaction

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/transaction"
	request "github.com/williamchang80/sea-apd/dto/request/transaction"
	transaction2 "github.com/williamchang80/sea-apd/mocks/repository/transaction"
	"reflect"
	"testing"
)

var (
	mockCreateTransactionRequest = request.TransactionRequest{
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		UserId:     "1",
	}
	mockTransactionEntity = transaction.Transaction{
		Status:     transactionStatus["ONPROGRESS"],
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		UserId:     "1",
	}
	mockUpdateTransaction = request.UpdateTransactionRequest{
		Id:     "1",
		Status: "accepted",
	}
	mockTransactionId = "1"
	mockUserId        = "1"
)

func TestNewTransactionUsecase(t *testing.T) {
	type args struct {
		repository transaction.TransactionRepository
	}
	tests := []struct {
		name string
		args args
		want transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: &TransactionUsecase{
				tr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToDomain(t *testing.T) {
	type args struct {
		productRequest request.TransactionRequest
	}
	tests := []struct {
		name string
		args args
		want transaction.Transaction
	}{
		{
			name: "success",
			args: args{
				productRequest: mockCreateTransactionRequest,
			},
			want: mockTransactionEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToDomain(tt.args.productRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionUsecase_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.TransactionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockCreateTransactionRequest,
			},
			wantErr: false,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				request: request.TransactionRequest{},
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.CreateTransaction(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionUsecase.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionUsecase_UpdateTransactionStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.UpdateTransactionRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockUpdateTransaction,
			},
			wantErr: false,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				request: request.UpdateTransactionRequest{
					Id:     "1",
					Status: "Fail",
				},
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			if err := c.UpdateTransactionStatus(tt.args.request); err != nil && !tt.wantErr {
				t.Errorf("TransactionUsecase.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionUsecase_GetTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		want     transaction.Transaction
		initMock func() transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockTransactionId,
			},
			wantErr: false,
			want:    mockTransactionEntity,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				request: "",
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			if p, err := c.GetTransactionById(tt.args.request); (err != nil || reflect.DeepEqual(p, tt.args)) && !tt.wantErr {
				t.Errorf("TransactionUsecase.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionUsecase_GetTransactionHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		want     []transaction.Transaction
		initMock func() transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockUserId,
			},
			wantErr: false,
			want:    []transaction.Transaction{},
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				request: "",
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				return NewTransactionUsecase(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			if p, err := c.GetTransactionHistory(tt.args.request); (err != nil || reflect.DeepEqual(p, tt.args)) && !tt.wantErr {
				t.Errorf("TransactionUsecase.GetTransactionHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
