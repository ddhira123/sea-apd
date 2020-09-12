package transaction

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/common/mailer"
	merchant3 "github.com/williamchang80/sea-apd/domain/merchant"
	product2 "github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/domain/transaction"
	request "github.com/williamchang80/sea-apd/dto/request/transaction"
	merchant4 "github.com/williamchang80/sea-apd/mocks/repository/merchant"
	transaction2 "github.com/williamchang80/sea-apd/mocks/repository/transaction"
	"github.com/williamchang80/sea-apd/mocks/usecase/merchant"
	"github.com/williamchang80/sea-apd/mocks/usecase/product"
	"github.com/williamchang80/sea-apd/mocks/usecase/user"
	merchant2 "github.com/williamchang80/sea-apd/usecase/merchant"
	"reflect"
	"testing"
)

var (
	mockCreateCartRequest = request.CreateCartRequest{
		UserId:     "1",
		MerchantId: "1",
	}
	mockConfirmationTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_CONFIRMATION),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockAcceptedTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.ACCEPTED),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockDeliveryTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_DELIVERY),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}

	mockPaymentRequest = request.PaymentRequest{
		CustomerId:    "1",
		BankNumber:    "123",
		BankName:      "123",
		TransactionId: "1",
	}
	mockUpdateTransaction = request.UpdateTransactionRequest{
		TransactionId: "1",
		Status:        transaction_status.ACCEPTED,
	}
	mockTransactionId = "1"
	mockUserId        = "1"
)

func TestNewTransactionUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		repository     transaction.TransactionRepository
		usecase        merchant3.MerchantUsecase
		productUsecase product2.ProductUsecase
	}
	tests := []struct {
		name string
		args args
		want transaction.TransactionUsecase
	}{
		{
			name: "success",
			args: args{
				repository:     nil,
				usecase:        merchant.NewMockUsecase(ctrl),
				productUsecase: product.NewMockUsecase(ctrl),
			},
			want: &TransactionUsecase{
				tr:              nil,
				merchantUseCase: merchant.NewMockUsecase(ctrl),
				productUseCase:  product.NewMockUsecase(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionUsecase(tt.args.repository, tt.args.usecase,
				tt.args.productUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToDomain(t *testing.T) {
	type args struct {
		productRequest request.CreateCartRequest
	}
	tests := []struct {
		name string
		args args
		want transaction.Transaction
	}{
		{
			name: "success",
			args: args{
				productRequest: mockCreateCartRequest,
			},
			want: mockConfirmationTransactionEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTransactionRequestToDomain(tt.args.productRequest);
				reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("ConvertToDomain() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestTransactionUsecase_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.CreateCartRequest
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
				request: mockCreateCartRequest,
			},
			wantErr: false,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				request: request.CreateCartRequest{},
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.CreateCart(tt.args.request)
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
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				request: request.UpdateTransactionRequest{
					TransactionId: "1",
					Status:        transaction_status.OTHER,
				},
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
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
			want:    mockAcceptedTransactionEntity,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
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
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
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
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
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
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
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

func TestTransactionUsecase_GetMerchantRequestItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		merchantId string
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
				merchantId: mockUserId,
			},
			wantErr: false,
			want:    []transaction.Transaction{},
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				merchantId: "",
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			if p, err := c.GetMerchantRequestItem(tt.args.merchantId);
				(err != nil || reflect.DeepEqual(p, tt.args)) && !tt.wantErr {
				t.Errorf("TransactionUsecase.GetMerchantRequestItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionUsecase_PayTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.PaymentRequest
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
				request: mockPaymentRequest,
			},
			wantErr: true,
			want:    []transaction.Transaction{},
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
		{
			name: "failed with unmatched status",
			args: args{
				request: request.PaymentRequest{},
			},
			wantErr: true,
			initMock: func() transaction.TransactionUsecase {
				t := transaction2.NewMockRepository(ctrl)
				u := merchant.NewMockUsecase(ctrl)
				p := product.NewMockUsecase(ctrl)
				return NewTransactionUsecase(t, u, p)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			if err := c.PayTransaction(tt.args.request);
				err != nil && !tt.wantErr {
				t.Errorf("TransactionUsecase.PayTransaction() error = %v, "+
					"wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNotifyAdminObserver_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mailer.InitMail()
	defer ctrl.Finish()
	type args struct {
		transaction transaction.Transaction
		status      transaction_status.TransactionStatus
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() merchant3.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				transaction: mockConfirmationTransactionEntity,
				status:      transaction_status.WAITING_CONFIRMATION,
			},
			initMock: func() merchant3.MerchantUsecase {
				repo := merchant4.NewMockRepository(ctrl)
				usecase := user.NewMockUsecase(ctrl)
				return merchant2.NewMerchantUsecase(repo, usecase)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obs := NotifyAdminObserver{}
			mock := tt.initMock()
			if err := obs.Update(tt.args.transaction, mock);
				err != nil && !tt.wantErr {
				t.Errorf("NotifyAdminObserver.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
