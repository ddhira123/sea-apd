package transfer

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/transfer"
	"github.com/williamchang80/sea-apd/dto/request/transfer"
	transfer_repository "github.com/williamchang80/sea-apd/mocks/repository/transfer"
	transfer_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/transfer"
	transfer_usecase "github.com/williamchang80/sea-apd/usecase/transfer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	mockId = "1"
	mockCreateTransferRequest = transfer.CreateTransferHistoryRequest{
		BankNumber: "Mock Number",
		BankName:   "Mock Name",
		Amount:     1000,
		MerchantId: "1",
	}
)

func TestNewTransferController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := transfer_repository.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}
	tests := []struct {
		name     string
		args     args
		want     domain.TransferController
		initMock func() domain.TransferUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &TransferController{
				usecase: transfer_usecase.NewTransferUsecase(repo, nil),
			},
			initMock: func() domain.TransferUsecase {
				return transfer_mock_usecase.NewMockUsecase(ctrl)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()

			if got := NewTransferController(tt.args.ctx, mock); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewTransferController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferController_GetTransferHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx       *echo.Echo
		getParams func() url.Values
	}

	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.TransferUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					q.Set("merchantId", mockId)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.TransferUsecase {
				c := transfer_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "failed with no params",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.TransferUsecase {
				c := transfer_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			req, err := http.NewRequest(echo.GET, "api/transfers"+"?"+tt.args.getParams().Encode(), strings.NewReader(""))
			if err != nil {
				t.Errorf("GetTransferHistory() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewTransferController(c, mock)
			if got := controller.GetTransferHistory(ctx); (got != nil) != tt.wantErr {
				t.Errorf("GetTransferHistory() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestTransferController_CreateTransferHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request transfer.CreateTransferHistoryRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.TransferUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockCreateTransferRequest,
			},
			wantErr: false,
			initMock: func() domain.TransferUsecase {
				c := transfer_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: transfer.CreateTransferHistoryRequest{},
			},
			wantErr: false,
			initMock: func() domain.TransferUsecase {
				c := transfer_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.POST, "api/product", strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("CreateTransferHistory() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewTransferController(c, mock)
			if got := controller.CreateTransferHistory(ctx); (got != nil) != tt.wantErr {
				t.Errorf("CreateTransferHistory() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}
